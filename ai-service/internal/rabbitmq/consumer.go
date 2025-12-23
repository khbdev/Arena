package rabbitmq

import (
	"fmt"
	"log"

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
		log.Fatalf("Queue’dan xabar ololmadi: %v", err)
	}

	fmt.Println("🐇 RabbitMQ consumer ishga tushdi...")

	
	pool := workerpool.New(5, 20, testService)
	pool.Start()


	for msg := range msgs {
		pool.Submit(msg.Body)
	}

	select {} 
}
