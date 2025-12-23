package config

import (
	"log"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

var (
	rabbitConn    *amqp.Connection
	rabbitChannel *amqp.Channel
	once          sync.Once
)

func InitRabbitMQ() {
	once.Do(func() {
		url := os.Getenv("RABBITMQ_URL")
		if url == "" {
			log.Fatal("RABBITMQ_URL env topilmadi")
		}

		conn, err := amqp.Dial(url)
		if err != nil {
			log.Fatalf("RabbitMQ connect error: %v", err)
		}

		ch, err := conn.Channel()
		if err != nil {
			log.Fatalf("RabbitMQ channel error: %v", err)
		}

		rabbitConn = conn
		rabbitChannel = ch

		log.Println("✅ RabbitMQ connected & channel opened")
	})
}

func RabbitChannel() *amqp.Channel {
	if rabbitChannel == nil {
		log.Fatal("RabbitMQ hali init qilinmagan")
	}
	return rabbitChannel
}


func CloseRabbitMQ() {
	if rabbitChannel != nil {
		_ = rabbitChannel.Close()
	}
	if rabbitConn != nil {
		_ = rabbitConn.Close()
	}
}
