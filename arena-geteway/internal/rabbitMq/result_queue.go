package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)


func PublishResultQueue(ch *amqp.Channel, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",            
		"result_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
