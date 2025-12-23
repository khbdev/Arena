package config

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func InitTelegramBot() *tgbotapi.BotAPI {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("❌ TELEGRAM_BOT_TOKEN env da topilmadi")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("❌ Telegram bot init error:", err)
	}

	log.Println("🤖 Telegram bot ishga tushdi:", bot.Self.UserName)

	return bot
}
