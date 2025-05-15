package structure

type CreateTariff struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateTariffAssignments struct {
	UserID   string `json:"user_id" validate:"required"`
	TariffID string `json:"tariff_id" validate:"required"`
}
