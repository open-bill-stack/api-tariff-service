package tariff

import (
	"api-tariff-service/internal/module/tariff/structure"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type HttpHandle struct {
	log     *zap.Logger
	service *Service
}

func NewHttpHandler(p Params) (HttpResult, error) {
	return HttpResult{
		Router: &HttpHandle{
			log:     p.Log,
			service: p.Service,
		},
	}, nil
}

func (h *HttpHandle) Register(app *fiber.App) {
	groupTariff := app.Group("/tariff")
	groupTariff.Post("", h.CreateTariff)
	groupTariff.Get("", h.ListTariffs)
	groupTariff.Get("/:id", h.GetTariff)
	groupTariff.Put("/:id", h.UpdateTariff)
	groupTariff.Delete("/:id", h.DeleteTariff)

}

func (h *HttpHandle) CreateTariff(c *fiber.Ctx) error {
	var req structure.CreateTariff
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())

	// Парсимо JSON з тіла запиту
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Перевіряємо валідацію структури
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	tariff := Tariff{
		Name: req.Name,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true,
		},
	}
	createTariff, err := h.service.CreateTariff(c.Context(), &tariff)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(structure.ResponseTariff{
		ID:          createTariff.ID.String(),
		Name:        createTariff.Name,
		Description: createTariff.Description.String,
	})
}
func (h *HttpHandle) GetTariff(c *fiber.Ctx) error {
	id := c.Params("id")
	var tariffID, errParse = uuid.Parse(id)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	tariff, err := h.service.GetTariffByID(c.Context(), tariffID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if tariff == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}
	return c.Status(fiber.StatusOK).JSON(structure.ResponseTariff{
		ID:          tariff.ID.String(),
		Name:        tariff.Name,
		Description: tariff.Description.String,
	})
}
func (h *HttpHandle) ListTariffs(c *fiber.Ctx) error {
	tariffs, err := h.service.ListTariffs(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var structTariffs = make([]structure.ResponseTariff, 0)
	for _, item := range tariffs {
		structTariffs = append(structTariffs, structure.ResponseTariff{
			ID:          item.ID.String(),
			Name:        item.Name,
			Description: item.Description.String,
		})
	}

	return c.Status(fiber.StatusOK).JSON(structTariffs)
}
func (h *HttpHandle) UpdateTariff(c *fiber.Ctx) error {
	id := c.Params("id")
	var tariffID, errParse = uuid.Parse(id)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	var req structure.CreateTariff
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())

	// Парсимо JSON з тіла запиту
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Перевіряємо валідацію структури
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	tariff := Tariff{
		ID: pgtype.UUID{
			Bytes: tariffID,
			Valid: true,
		},
		Name: req.Name,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true,
		},
	}
	createTariff, err := h.service.UpdateTariffByID(c.Context(), &tariff)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(structure.ResponseTariff{
		ID:          createTariff.ID.String(),
		Name:        createTariff.Name,
		Description: createTariff.Description.String,
	})
}
func (h *HttpHandle) DeleteTariff(c *fiber.Ctx) error {
	id := c.Params("id")
	var tariffID, errParse = uuid.Parse(id)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	status, err := h.service.DeleteTariffByID(c.Context(), tariffID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if !status {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Tariff not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tariff deleted successfully",
	})
}
