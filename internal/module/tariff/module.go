package tariff

import "go.uber.org/fx"

var Module = fx.Module("tariff",
	fx.Provide(NewHttpHandler),
	fx.Provide(NewGrpcHandler),

	fx.Provide(NewService),
	fx.Provide(NewRepository),
	fx.Provide(NewEventPublisher),

	fx.Invoke(RunEvent),
)
