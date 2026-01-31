package event

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// UnsuccessEvent represents a failed operation event payload
// that will be published as JSON.
type UnsuccessEvent struct {
	TelegramID int64  `json:"telegram_id"`
	Message    string `json:"message"`
	ErrorMsg   string `json:"error_msg"`
}

// PublishUnsuccess publishes an unsuccessful event to the unsuccess_queue
// with the provided telegramID and error message.
func PublishUnsuccess(ch *amqp.Channel, telegramID int64, errorMsg string) error {
	event := UnsuccessEvent{
		TelegramID: telegramID,
		Message:    "An error occurred while generating the test. Please try again later.",
		ErrorMsg:   errorMsg,
	}

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal unsuccess event to JSON: %v", err)
	}

	queueName := "unsuccess_queue"
	routingKey := "unsuccess_key"
	exchangeName := "direct_exchange"

	err = ch.Publish(
		exchangeName, // exchange name
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish unsuccess event to RabbitMQ: %v", err)
	}

	log.Println("Event successfully published to", queueName+":", string(body))
	return nil
}
