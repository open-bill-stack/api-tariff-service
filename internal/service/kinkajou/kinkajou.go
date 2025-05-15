package kinkajou

import (
	"api-tariff-service/internal/service/config"
	"api-tariff-service/internal/service/kinkajou/router"
	"context"
	"github.com/open-bill-stack/kinkajou"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	EventApp    *kinkajou.App
	AQMPChannel *amqp.Channel
	Log         *zap.Logger
	Config      *config.Config
	Router      []router.Router `group:"eventRoutes"`
}
type Result struct {
	fx.Out
	EventApp *kinkajou.App
}

func NewEventApp() (Result, error) {
	return Result{
		EventApp: kinkajou.New(),
	}, nil
}

func RunEventApp(lc fx.Lifecycle, p Params) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			for _, r := range p.Router {
				r.Register(p.EventApp)
			}

			ch, err := p.AQMPChannel.Consume(
				"tariff.service",
				"",
				false, // autoAck
				false, // exclusive
				false, // noLocal
				false, // noWait
				nil,   // args
			)
			if err != nil {
				return err
			}
			go func() {
				if err := p.EventApp.Listen(ch); err != nil {
					p.Log.Panic("Error starting Fiber server:", zap.Error(err))
				}
			}()
			return nil
		},
	})

}

var Module = fx.Module(
	"FiberAppModule",
	fx.Provide(
		NewEventApp,
	),
	fx.Invoke(
		RunEventApp,
	),
)
