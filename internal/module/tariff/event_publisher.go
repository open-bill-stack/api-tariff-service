package tariff

import (
	"context"
	"github.com/google/uuid"
)

type EventPublisher interface {
	PublishTariffDelete(ctx context.Context, id uuid.UUID) error
}
