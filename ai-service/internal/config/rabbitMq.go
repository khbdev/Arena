package config

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

type RabbitMQConnection struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ() *RabbitMQConnection {
	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASS")
	host := os.Getenv("RABBITMQ_HOST")

	url := fmt.Sprintf("amqp://%s:%s@%s/%%2F", user, pass, host)

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	fmt.Println("Successfully connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open RabbitMQ channel: %v", err)
	}
	fmt.Println("RabbitMQ channel opened successfully")

	r := &RabbitMQConnection{
		Conn:    conn,
		Channel: ch,
	}

	if err := r.setupQueues(); err != nil {
		log.Fatalf("RabbitMQ queue setup failed: %v", err)
	}
	fmt.Println("RabbitMQ setup completed successfully")

	return r
}

func (r *RabbitMQConnection) setupQueues() error {
	exchange := "direct_exchange"

	// --- Arena Queue configuration ---
	arenaQueue := "arena_queue"
	arenaRoutingKey := "queue_key"
	arenaDLQ := "arena_queue_dlq"
	arenaRetry := "arena_queue_retry"

	// --- Notifications Queue configuration ---
	notificationsQueue := "notifications_queue"
	notificationsRoutingKey := "notif_key"
	notificationsRetry := "notifications_queue_retry"

	// --- Unsuccessful Queue configuration ---
	unsuccessQueue := "unsuccess_queue"
	unsuccessRoutingKey := "unsuccess_key"
	unsuccessDLQ := "unsuccess_queue_dlq"
	unsuccessRetry := "unsuccess_queue_retry"

	// --- Exchange ---
	if err := r.Channel.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// --- Arena Dead Letter Queue (DLQ) ---
	if _, err := r.Channel.QueueDeclare(
		arenaDLQ,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("failed to declare arena DLQ: %w", err)
	}

	// --- Arena Retry Queue (10s TTL) ---
	if _, err := r.Channel.QueueDeclare(
		arenaRetry,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    exchange,
			"x-dead-letter-routing-key": arenaRoutingKey,
			"x-message-ttl":             int32(10000),
		},
	); err != nil {
		return fmt.Errorf("failed to declare arena retry queue: %w", err)
	}

	// --- Arena Main Queue (priority + DLQ) ---
	if _, err := r.Channel.QueueDeclare(
		arenaQueue,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    exchange,
			"x-dead-letter-routing-key": arenaDLQ,
			"x-max-priority":            int32(10),
		},
	); err != nil {
		return fmt.Errorf("failed to declare arena queue: %w", err)
	}

	if err := r.Channel.QueueBind(arenaQueue, arenaRoutingKey, exchange, false, nil); err != nil {
		return fmt.Errorf("failed to bind arena queue: %w", err)
	}

	// --- Notifications Retry Queue (1 hour TTL) ---
	if _, err := r.Channel.QueueDeclare(
		notificationsRetry,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    exchange,
			"x-dead-letter-routing-key": notificationsRoutingKey,
			"x-message-ttl":             int32(3600 * 1000),
		},
	); err != nil {
		return fmt.Errorf("failed to declare notifications retry queue: %w", err)
	}

	// --- Notifications Main Queue ---
	if _, err := r.Channel.QueueDeclare(
		notificationsQueue,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    exchange,
			"x-dead-letter-routing-key": notificationsRetry,
		},
	); err != nil {
		return fmt.Errorf("failed to declare notifications queue: %w", err)
	}

	if err := r.Channel.QueueBind(notificationsQueue, notificationsRoutingKey, exchange, false, nil); err != nil {
		return fmt.Errorf("failed to bind notifications queue: %w", err)
	}

	// --- Unsuccessful Dead Letter Queue (DLQ) ---
	if _, err := r.Channel.QueueDeclare(
		unsuccessDLQ,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("failed to declare unsuccessful DLQ: %w", err)
	}

	// --- Unsuccessful Retry Queue (30s TTL) ---
	if _, err := r.Channel.QueueDeclare(
		unsuccessRetry,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    exchange,
			"x-dead-letter-routing-key": unsuccessRoutingKey,
			"x-message-ttl":             int32(30000),
		},
	); err != nil {
		return fmt.Errorf("failed to declare unsuccessful retry queue: %w", err)
	}

	// --- Unsuccessful Main Queue (priority + DLQ) ---
	if _, err := r.Channel.QueueDeclare(
		unsuccessQueue,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    exchange,
			"x-dead-letter-routing-key": unsuccessDLQ,
			"x-max-priority":            int32(10),
		},
	); err != nil {
		return fmt.Errorf("failed to declare unsuccessful queue: %w", err)
	}

	if err := r.Channel.QueueBind(unsuccessQueue, unsuccessRoutingKey, exchange, false, nil); err != nil {
		return fmt.Errorf("failed to bind unsuccessful queue: %w", err)
	}

	return nil
}
