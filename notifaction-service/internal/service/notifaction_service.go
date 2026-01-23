package service

import (
	"fmt"
	"notifaction-service/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendTestNotification(telegramID int64, testID string) {
	bot := config.TelegramBot()

	msgText := fmt.Sprintf(
		"📝 Yangi test tayyor!\n\n"+
			"🆔 Test ID: `%s`\n\n"+
			"⏰ Testni o‘z vaqtida topshirishni unutmang\n"+
			"🚀 Omad tilaymiz!\n\n"+
			"🔗 Bot link: @arena_rep_bot",
		testID,
	)

	msg := tgbotapi.NewMessage(telegramID, msgText)
	msg.ParseMode = "Markdown"

	_, _ = bot.Send(msg) 
}
