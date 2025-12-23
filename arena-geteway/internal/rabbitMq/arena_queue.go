package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)


func PublishArenaQueue(ch *amqp.Channel, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",            
		"arena_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
