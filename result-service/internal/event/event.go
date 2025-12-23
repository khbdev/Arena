package event

import (
	"encoding/json"
	"fmt"
	"log"

	"result-service/internal/config"
	"result-service/internal/model"

	"github.com/streadway/amqp"
)

func publish(queue string, body []byte) {
	err := config.RMQ.Channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
	if err != nil {
		log.Println(" Event publish error:", err)
	}
	fmt.Println("Event Publishid")
}

func PublishErrorEvent(e model.ErrorEvent) {
	b, _ := json.Marshal(e)
	publish(config.UserStatikQueue, b)
}

func PublishUserResultEvent(e model.UserResultEvent) {
	b, _ := json.Marshal(e)
	publish(config.UserStatikQueue, b)
}
