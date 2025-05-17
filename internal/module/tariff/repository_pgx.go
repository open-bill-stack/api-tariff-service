package tariff

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type sqlRepo struct {
	db *pgxpool.Pool
}

func (s *sqlRepo) ExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM tariffs WHERE id = $1)`
	row := s.db.QueryRow(ctx, query, id)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if tariff exists: %w", err)
	}
	return exists, nil
}
func (s *sqlRepo) CreateTariff(ctx context.Context, item *Tariff) (*Tariff, error) {
	query := `INSERT INTO tariffs (name, description) VALUES ($1, $2) RETURNING id, name, description`
	row := s.db.QueryRow(ctx, query, item.Name, item.Description)
	var result Tariff
	err := row.Scan(&result.ID, &result.Name, &result.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to create tariff: %w", err)
	}

	return &result, nil
}

func (s *sqlRepo) GetTariffByID(ctx context.Context, id uuid.UUID) (*Tariff, error) {
	query := `SELECT * FROM tariffs WHERE id = $1`
	row := s.db.QueryRow(ctx, query, id)
	var result Tariff
	err := row.Scan(&result.ID, &result.Name, &result.Description)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get tariff by ID: %w", err)
	}
	return &result, nil
}

func (s *sqlRepo) UpdateTariffByID(ctx context.Context, item *Tariff) (*Tariff, error) {
	query := `UPDATE tariffs SET name = $1, description = $2 WHERE id = $3 RETURNING id, name, description`
	row := s.db.QueryRow(ctx, query, item.Name, item.Description, item.ID)
	var result Tariff
	err := row.Scan(&result.ID, &result.Name, &result.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to update tariff: %w", err)
	}

	return &result, nil
}

func (s *sqlRepo) DeleteTariffByID(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `DELETE FROM tariffs WHERE id = $1`
	cmdTag, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete tariff: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return false, nil
	}
	return true, nil
}

func (s *sqlRepo) ListTariffs(ctx context.Context) ([]Tariff, error) {
	query := `SELECT * FROM tariffs`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list tariffs: %w", err)
	}

	var tariffs []Tariff
	for rows.Next() {
		var tariff Tariff
		if err := rows.Scan(&tariff.ID, &tariff.Name, &tariff.Description); err != nil {
			return nil, fmt.Errorf("failed to scan tariff: %w", err)
		}
		tariffs = append(tariffs, tariff)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return tariffs, nil
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &sqlRepo{db}
}
