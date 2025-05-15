package tariff

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type Service struct {
	repo           Repository
	userGrpcClient UserGrpcClient
}

func NewService(r Repository, client UserGrpcClient) *Service {
	return &Service{
		repo:           r,
		userGrpcClient: client,
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

func (s *Service) CreateTariffAssignments(ctx context.Context, item *Assignment) (*Assignment, error) {
	status, err := s.userGrpcClient.UserExistsByID(ctx, item.UserID.Bytes)
	if err != nil {
		return nil, err
	}
	if !status {
		return nil, fmt.Errorf("userID not exist")
	}

	return s.repo.CreateTariffAssignments(ctx, item)
}

func (s *Service) GetTariffAssignmentsByID(ctx context.Context, id uuid.UUID) (*Assignment, error) {
	return s.repo.GetTariffAssignmentsByID(ctx, id)
}

func (s *Service) UpdateTariffAssignmentsByID(ctx context.Context, item *Assignment) (*Assignment, error) {
	return s.repo.UpdateTariffAssignmentsByID(ctx, item)
}

func (s *Service) DeleteTariffAssignmentsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.repo.DeleteTariffAssignmentsByID(ctx, id)
}

func (s *Service) ListTariffAssignments(ctx context.Context) ([]Assignment, error) {
	return s.repo.ListTariffAssignments(ctx)
}

func (s *Service) DeleteTariffAssignmentsByUserID(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.repo.DeleteTariffAssignmentsByUserID(ctx, id)
}
