// internal/bot/bot.go
package bot

import (
	"log"

	"geteway/internal/client"
	"geteway/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/streadway/amqp"
)

type Bot struct {
	api           *tgbotapi.BotAPI
	userService   *service.UserService
	resultService *service.ResultService
	telegramClient *client.TelegramClient
	rabbitCh      *amqp.Channel
	userStates    map[int64]*UserState
}

func New(
	api *tgbotapi.BotAPI,
	userService *service.UserService,
	resultService *service.ResultService,
	rabbitCh *amqp.Channel,
	telegramClient *client.TelegramClient,
) *Bot {
	return &Bot{
		api:            api,
		userService:    userService,
		resultService:  resultService,
		rabbitCh:       rabbitCh,
		telegramClient: telegramClient,
		userStates:     make(map[int64]*UserState),
	}
}


func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)
	log.Println("🚀 Bot ishga tushdi")

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				b.handleCommand(update)
			} else {
				b.handleText(update)
			}
		} else if update.CallbackQuery != nil {
			b.handleCallback(update)
		}
	}
}