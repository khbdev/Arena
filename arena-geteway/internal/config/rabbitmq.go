package config

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}


func InitRabbitMQ() *RabbitMQ {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		log.Fatal("RABBITMQ_URL env da topilmadi")
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal("RabbitMQ connection error:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("RabbitMQ channel error:", err)
	}

	log.Println("RabbitMQ connected")

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}
}
