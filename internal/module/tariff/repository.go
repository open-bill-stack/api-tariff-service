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

	CreateTariffAssignments(ctx context.Context, item *Assignment) (*Assignment, error)
	GetTariffAssignmentsByID(ctx context.Context, id uuid.UUID) (*Assignment, error)
	UpdateTariffAssignmentsByID(ctx context.Context, item *Assignment) (*Assignment, error)
	DeleteTariffAssignmentsByID(ctx context.Context, id uuid.UUID) (bool, error)
	ListTariffAssignments(ctx context.Context) ([]Assignment, error)

	DeleteTariffAssignmentsByUserID(ctx context.Context, id uuid.UUID) (bool, error)
}
