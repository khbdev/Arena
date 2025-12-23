package config

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

var RMQ *RabbitMQ

const (
	ResultQueue     = "result_queue"
	UserStatikQueue = "user_statik_queue" 
	DLQQueue        = "dlq_result_queue"
)

const (
	MaxRetryCount = 3
	RetryDelay    = 5 * time.Second
)

func InitRabbitMQ(amqpURL string) (*RabbitMQ, error) {
	if RMQ != nil {
		return RMQ, nil
	}

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %w", err)
	}

	// result_queue
	_, err = ch.QueueDeclare(
		ResultQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare result queue: %w", err)
	}

	// user_statik_queue 🔥
	_, err = ch.QueueDeclare(
		UserStatikQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare user_statik_queue: %w", err)
	}

	// dlq_result_queue
	_, err = ch.QueueDeclare(
		DLQQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare DLQ queue: %w", err)
	}

	RMQ = &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}

	return RMQ, nil
}
