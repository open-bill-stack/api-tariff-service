package grpc

import (
	"api-tariff-service/internal/service/config"
	pb "api-tariff-service/internal/service/grpc/proto/user/v1"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Params struct {
	fx.In

	Log    *zap.Logger
	Config *config.Config
}

type Result struct {
	fx.Out
	GrpcClientConnApp     *grpc.ClientConn
	GrpcUserServiceClient *pb.UserServiceClient
}

func NewGrpcClientApp(p Params) (Result, error) {
	conn, err := grpc.NewClient(
		p.Config.Service.UserGrpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return Result{}, fmt.Errorf("create grpc client: %w", err)
	}
	client := pb.NewUserServiceClient(conn)

	return Result{
		GrpcClientConnApp:     conn,
		GrpcUserServiceClient: &client,
	}, nil
}

var Module = fx.Module(
	"GrpcAppModule",
	fx.Provide(
		NewGrpcClientApp,
	),
)
