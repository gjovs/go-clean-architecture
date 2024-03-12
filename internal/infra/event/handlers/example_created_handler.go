package handlers

import (
	"encoding/json"
	"sync"

	"github.com/gjovs/clean/pkg/events"
	"github.com/streadway/amqp"
)

type ExampleCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewExampleCreatedHandler(rabbitMQChannel *amqp.Channel) *ExampleCreatedHandler {
	return &ExampleCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (handler *ExampleCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	json, _ := json.Marshal(event.GetPayload())

	json_message_to_rabbit_mq := amqp.Publishing{
		ContentType: "application/json",
		Body:        json,
	}

	handler.RabbitMQChannel.Publish("amqp.direct", "", false, false, json_message_to_rabbit_mq)
}
