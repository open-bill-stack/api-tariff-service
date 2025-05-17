package tariff

import (
	"api-tariff-service/internal/module/tariff/structure"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitEventPublisher struct {
	ch *amqp.Channel
}

func NewEventPublisher(ch *amqp.Channel) EventPublisher {
	return &rabbitEventPublisher{ch}
}

func (r *rabbitEventPublisher) PublishTariffDelete(ctx context.Context, id uuid.UUID) error {
	var event = structure.Event{
		TariffUUID: id.String(),
	}

	marshal, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return r.ch.PublishWithContext(ctx,
		"tariff.events",  // exchange
		"tariff.deleted", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        marshal,
		})
}
