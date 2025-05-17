package tariff

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	CreateTariff(ctx context.Context, item *Tariff) (*Tariff, error)
	GetTariffByID(ctx context.Context, id uuid.UUID) (*Tariff, error)
	UpdateTariffByID(ctx context.Context, item *Tariff) (*Tariff, error)
	DeleteTariffByID(ctx context.Context, id uuid.UUID) (bool, error)
	ListTariffs(ctx context.Context) ([]Tariff, error)
	ExistsByID(ctx context.Context, id uuid.UUID) (bool, error)
}
