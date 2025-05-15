package tariff

import (
	pb "api-tariff-service/internal/service/grpc/proto/user/v1"
	"context"
	"github.com/google/uuid"
)

type userGrpcClient struct {
	client pb.UserServiceClient
}

func NewUserGrpcClient(client *pb.UserServiceClient) UserGrpcClient {
	return &userGrpcClient{*client}
}

func (u *userGrpcClient) UserExistsByID(ctx context.Context, userID uuid.UUID) (bool, error) {
	r, err := u.client.ExistsByID(ctx, &pb.ExistsByIDRequest{UserId: userID.String()})
	if err != nil {
		return false, err
	}
	return r.GetExists(), nil
}
