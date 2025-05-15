package tariff

import (
	"context"
	"github.com/google/uuid"
)

type UserGrpcClient interface {
	UserExistsByID(ctx context.Context, userID uuid.UUID) (bool, error)
}
