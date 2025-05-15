package structure

type ResponseTariff struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type ResponseTariffAssignments struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	TariffID string `json:"tariff_id"`
}
