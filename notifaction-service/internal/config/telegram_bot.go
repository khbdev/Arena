package config

import (
	"log"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	bot     *tgbotapi.BotAPI
	onceBot sync.Once
)


func InitTelegramBot() {
	onceBot.Do(func() {
		token := os.Getenv("TELEGRAM_BOT_TOKEN")
		if token == "" {
			log.Fatal("TELEGRAM_BOT_TOKEN env topilmadi")
		}

		var err error
		bot, err = tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Fatalf("Telegram bot init error: %v", err)
		}

		log.Printf("Telegram bot authorized on account %s", bot.Self.UserName)
	})
}


func TelegramBot() *tgbotapi.BotAPI {
	if bot == nil {
		log.Fatal("Telegram bot hali init qilinmagan, InitTelegramBot() chaqiring")
	}
	return bot
}
