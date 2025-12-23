package main

import (
	"log"
	"os"
	"result-service/internal/config"
	"result-service/internal/handler"

	"github.com/joho/godotenv"
)


func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	config.InitRedis()

		amqpURL := os.Getenv("RABBITMQ_URL")
	if amqpURL == "" {
		log.Fatal("RABBITMQ_URL not set in .env")
	}
	rmq, err := config.InitRabbitMQ(amqpURL)
	if err != nil {
		log.Fatalf("RabbitMQ init error: %v", err)
	}

	defer rmq.Conn.Close()
	defer rmq.Channel.Close()

	log.Println(" RabbitMQ connection established and queues declared!")

		handler.StartConsumer(rmq)
	
}
