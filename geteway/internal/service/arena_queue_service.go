package service

import (
	"geteway/internal/model"
	rabbitmq "geteway/internal/rabbitMq"

	"github.com/streadway/amqp"
)


func HandleArenaRequest(ch *amqp.Channel, telegramID int64, prompt string, count int) error {
	data := model.ArenaRequest{
		TelegramID: telegramID,
		Prompt:     prompt,
		Count:      count,
	}

	return rabbitmq.PublishArenaQueue(ch, data)
}
