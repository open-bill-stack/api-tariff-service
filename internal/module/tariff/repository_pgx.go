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

func (s *sqlRepo) DeleteTariffAssignmentsByUserID(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `DELETE FROM tariff_assignments WHERE user_id = $1`
	cmdTag, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return false, err
	}
	if cmdTag.RowsAffected() == 0 {
		return false, nil // або поверни спеціальну помилку NotFound, якщо треба
	}
	return true, nil
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &sqlRepo{db}
}

func (s *sqlRepo) CreateTariff(ctx context.Context, item *Tariff) (*Tariff, error) {
	query := `INSERT INTO tariffs (name, description) VALUES ($1, $2) RETURNING *`
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
	query := `UPDATE tariffs SET name = $1, description = $2 WHERE id = $3 RETURNING *`
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

func (s *sqlRepo) CreateTariffAssignments(ctx context.Context, item *Assignment) (*Assignment, error) {
	query := `INSERT INTO tariff_assignments (user_id, tariff_id) VALUES ($1, $2) RETURNING *`
	row := s.db.QueryRow(ctx, query, item.UserID, item.TariffID)
	var result Assignment
	err := row.Scan(&result.ID, &result.UserID, &result.TariffID)
	if err != nil {
		return nil, fmt.Errorf("failed to create tariff assignment: %w", err)
	}

	return &result, nil
}

func (s *sqlRepo) GetTariffAssignmentsByID(ctx context.Context, id uuid.UUID) (*Assignment, error) {
	query := `SELECT * FROM tariff_assignments WHERE id = $1`
	row := s.db.QueryRow(ctx, query, id)
	var result Assignment
	err := row.Scan(&result.ID, &result.UserID, &result.TariffID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get tariff assignment by ID: %w", err)
	}
	return &result, nil
}

func (s *sqlRepo) UpdateTariffAssignmentsByID(ctx context.Context, item *Assignment) (*Assignment, error) {
	query := `UPDATE tariff_assignments SET user_id = $1, tariff_id = $2 WHERE id = $3 RETURNING *`
	row := s.db.QueryRow(ctx, query, item.UserID, item.TariffID, item.ID)
	var result Assignment
	err := row.Scan(&result.ID, &result.UserID, &result.TariffID)
	if err != nil {
		return nil, fmt.Errorf("failed to update tariff assignment: %w", err)
	}

	return &result, nil
}

func (s *sqlRepo) DeleteTariffAssignmentsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `DELETE FROM tariff_assignments WHERE id = $1`
	cmdTag, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete tariff assignment: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return false, nil
	}
	return true, nil
}

func (s *sqlRepo) ListTariffAssignments(ctx context.Context) ([]Assignment, error) {
	query := `SELECT * FROM tariff_assignments`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list tariff assignments: %w", err)
	}

	var assignments []Assignment
	for rows.Next() {
		var assignment Assignment
		if err := rows.Scan(&assignment.ID, &assignment.UserID, &assignment.TariffID); err != nil {
			return nil, fmt.Errorf("failed to scan tariff assignment: %w", err)
		}
		assignments = append(assignments, assignment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return assignments, nil
}
