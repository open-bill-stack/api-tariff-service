package tariff

import (
	httpRouter "api-tariff-service/internal/service/fiber/router"
	eventRouter "api-tariff-service/internal/service/kinkajou/router"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Log     *zap.Logger
	Service *Service
}
type HttpResult struct {
	fx.Out

	Router httpRouter.Router `group:"httpRoutes"`
}
type EventResult struct {
	fx.Out

	Router eventRouter.Router `group:"eventRoutes"`
}
