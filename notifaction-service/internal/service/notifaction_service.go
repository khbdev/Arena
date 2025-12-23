package service

import (
	"fmt"
	"log"
	"notifaction-service/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendTestNotification(telegramID int64, testID string) {
	bot := config.TelegramBot()

	// Oddiy link matni bilan xabar
	msgText := fmt.Sprintf(
		"📝 Yangi test tayyor!\n\n"+
			"🆔 Test ID: %s\n\n"+
			"⏰ Testni o‘z vaqtida topshirishni unutmang\n"+
			"🚀 Omad tilaymiz!\n\n"+
			"🔗 Bot link: @arena_rep_bot",
		testID,
	)

	msg := tgbotapi.NewMessage(telegramID, msgText)

	if _, err := bot.Send(msg); err != nil {
		log.Printf("❌ Telegram send error for %d: %v", telegramID, err)
	} else {
		log.Printf("✅ Test notification sent to %d (TestID: %s)", telegramID, testID)
	}
}
