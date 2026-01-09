package rabbitmq

import (
	"ai-service/internal/service"
	"ai-service/internal/workerpool"

	"github.com/streadway/amqp"
)

const queueName = "arena_queue"

func StartConsumer(ch *amqp.Channel, testService *service.TestService) {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	pool := workerpool.New(5, 20, testService)
	pool.Start()

	for msg := range msgs {
		pool.Submit(msg.Body)
	}

	select {}
}
