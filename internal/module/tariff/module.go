package tariff

import "go.uber.org/fx"

var Module = fx.Module("tariff",
	fx.Provide(NewEventHandler),
	fx.Provide(NewHttpHandler),

	fx.Provide(NewService),
	fx.Provide(NewRepository),
	fx.Provide(NewUserGrpcClient),

	fx.Invoke(RunEvent),
)
