package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"result-service/internal/config"
)

// Model struct shu fayl ichida
type Answer struct {
	QuestionID    string `json:"question_id"`
	CorrectAnswer string `json:"correct_answer"`
}

type TestResult struct {
	TelegramID    int64    `json:"telegram_id"`
	TestID        string   `json:"test_id"`
	Answers       []Answer `json:"answers"`
	UserFirstName string   `json:"user_first_name"`
}

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	amqpURL := os.Getenv("RABBITMQ_URL")
	if amqpURL == "" {
		log.Fatal("RABBITMQ_URL not set in .env")
	}

	// Init RabbitMQ
	rmq, err := config.InitRabbitMQ(amqpURL)
	if err != nil {
		log.Fatalf("RabbitMQ init error: %v", err)
	}
	defer rmq.Conn.Close()
	defer rmq.Channel.Close()

	log.Println("✅ RabbitMQ connection established!")

	// Model bilan JSON yaratish
		req := TestResult{
		TelegramID: 8170926102,
		TestID:     "TST-AGoLpP",
		Answers: []Answer{
			{
				QuestionID:    "q1_present_simple",
				CorrectAnswer: "B",
			},
		},
		UserFirstName: "Gopher",
	}


	// Marshal to JSON
	body, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	// Publish to RabbitMQ
	err = rmq.Channel.Publish(
		"",                  // exchange
		config.ResultQueue,  // queue name
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	log.Println("✅ Message published to result_queue")
}
