package rabbitmq

import (
	"ai-service/internal/domain"
	"log"

	"github.com/streadway/amqp"
)

const queueName = "arena_queue"

type Consumer struct {
	ch          *amqp.Channel
	testService domain.MessageProcessor
}

func NewConsumer(ch *amqp.Channel, 	testService domain.MessageProcessor) *Consumer {
	return &Consumer{
		ch:          ch,

	testService: testService,
	}
}

func (c *Consumer) Start() {
	msgs, err := c.ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("consume error:", err)
		return
	}

	for msg := range msgs {

	for msg := range msgs {
    c.testService.ProcessMessage(msg.Body)
    _ = msg.Ack(false)
}

		_ = msg.Ack(false)
	}
}