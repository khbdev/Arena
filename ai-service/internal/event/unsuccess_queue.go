package event

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// UnsuccessEvent struct xabarni JSON ga o‘girish uchun
type UnsuccessEvent struct {
	TelegramID int64  `json:"telegram_id"`
	Message string `json:"message"`
	ErrorMsg   string `json:"error_msg"`
}

// PublishUnsuccess telegramID va errorMsg ni unsuccess_queue ga yuboradi
func PublishUnsuccess(ch *amqp.Channel, telegramID int64, errorMsg string) error {
	event := UnsuccessEvent{
		TelegramID: telegramID,
		Message: "Test Yaratishda Xatolik Yuz berdi iltimos keyinroq urinib koring",
		ErrorMsg:   errorMsg,
	}

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("JSON marshal xatolik: %v", err)
	}

	queueName := "unsuccess_queue"
	routingKey := "unsuccess_key"
	exchangeName := "direct_exchange"

	err = ch.Publish(
		exchangeName, // exchange nomi
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("RabbitMQ publish xatolik: %v", err)
	}

	log.Println("Event", queueName, "ga yuborildi:", string(body))
	return nil
}
