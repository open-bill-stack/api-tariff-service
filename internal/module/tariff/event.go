package tariff

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ParamsRun struct {
	fx.In

	Log         *zap.Logger
	AQMPChannel *amqp.Channel
}

func RunEvent(lc fx.Lifecycle, p ParamsRun) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			_, err := p.AQMPChannel.QueueDeclare(
				"tariff.service", // name
				false,            // durable
				false,            // delete when unused
				false,            // exclusive
				false,            // no-wait
				nil,              // arguments
			)
			if err != nil {
				return err
			}
			return p.AQMPChannel.QueueBind(
				"tariff.service",
				"user.*",      // routing key
				"user.events", // name
				false,
				nil,
			)
		},
	})
}
