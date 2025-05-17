package app

import (
	"api-tariff-service/internal/module/tariff"
	"api-tariff-service/internal/service/config"
	"api-tariff-service/internal/service/database"
	"api-tariff-service/internal/service/fiber"
	"api-tariff-service/internal/service/fiber/middleware"
	"api-tariff-service/internal/service/grpc"
	"api-tariff-service/internal/service/jwt"
	"api-tariff-service/internal/service/logger"
	"api-tariff-service/internal/service/rabbitmq"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func Run(cmd *cobra.Command) {
	fx.New(
		fx.Provide(func() *cobra.Command { return cmd }),

		logger.Module,
		config.Module,
		database.Module,
		jwt.Module,

		// global middleware
		middleware.Module,

		// fiber http
		tariff.Module,

		rabbitmq.Module,
		// fiber
		fiber.Module,

		// grpc
		grpc.ModuleServer,
	).Run()
}
