package tariff

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Tariff struct {
	ID          pgtype.UUID `db:"id"` // UUID асоціації
	Name        string      `db:"name"`
	Description pgtype.Text `db:"description"`
}

type Resource struct {
	ID       pgtype.UUID `db:"id"`        // UUID асоціації
	TariffID pgtype.UUID `db:"tariff_id"` // UUID тарифу
}

type Assignment struct {
	ID       pgtype.UUID `db:"id"`        // UUID асоціації
	UserID   pgtype.UUID `db:"user_id"`   // UUID користувача
	TariffID pgtype.UUID `db:"tariff_id"` // UUID тарифу
}
