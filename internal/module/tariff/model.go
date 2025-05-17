package tariff

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Tariff struct {
	ID          pgtype.UUID `db:"id"` // UUID асоціації
	Name        string      `db:"name"`
	Description pgtype.Text `db:"description"`
}
