package tariff

import (
	"context"
	"github.com/google/uuid"
)

type Service struct {
	repo           Repository
	eventPublisher EventPublisher
}

func NewService(r Repository, e EventPublisher) *Service {
	return &Service{
		repo:           r,
		eventPublisher: e,
	}
}

func (s *Service) CreateTariff(ctx context.Context, item *Tariff) (*Tariff, error) {
	return s.repo.CreateTariff(ctx, item)
}

func (s *Service) GetTariffByID(ctx context.Context, id uuid.UUID) (*Tariff, error) {
	return s.repo.GetTariffByID(ctx, id)
}

func (s *Service) UpdateTariffByID(ctx context.Context, item *Tariff) (*Tariff, error) {
	return s.repo.UpdateTariffByID(ctx, item)
}

func (s *Service) DeleteTariffByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.repo.DeleteTariffByID(ctx, id)
}

func (s *Service) ListTariffs(ctx context.Context) ([]Tariff, error) {
	return s.repo.ListTariffs(ctx)
}

func (s *Service) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.repo.ExistsByID(ctx, id)
}
