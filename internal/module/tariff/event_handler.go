package tariff

import (
	"api-tariff-service/internal/module/tariff/structure"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/open-bill-stack/kinkajou"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type EventHandle struct {
	log     *zap.Logger
	service *Service
}

func NewEventHandler(p Params) (EventResult, error) {
	return EventResult{
		Router: &EventHandle{
			log:     p.Log,
			service: p.Service,
		},
	}, nil
}

func (h *EventHandle) Register(app *kinkajou.App) {
	app.OnEvent("user.deleted", h.UserDeleted)
}

func (h *EventHandle) UserDeleted(message *amqp.Delivery) {
	var event structure.Event
	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		_ = message.Reject(true)
	}

	var userID, errParse = uuid.Parse(event.UserUUID)
	if errParse != nil {
		_ = message.Reject(true)
	}

	if _, err := h.service.DeleteTariffAssignmentsByUserID(context.Background(), userID); err != nil {
		_ = message.Reject(true)
	}
	_ = message.Ack(false)

}
