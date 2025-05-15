package middleware

import (
	"api-tariff-service/internal/service/fiber/middleware/healthcheck"
	"api-tariff-service/internal/service/fiber/middleware/jwt"
	middlewareRecover "api-tariff-service/internal/service/fiber/middleware/recover"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"MiddlewareModule",
	fx.Provide(
		healthcheck.NewMiddleware,
		middlewareRecover.NewMiddleware,
		jwt.NewService,
	),
)
