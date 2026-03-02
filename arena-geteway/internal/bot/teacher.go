// internal/bot/teacher.go
package bot

import (
	"log"
	"strconv"

	rabbitmq "geteway/internal/rabbitMq"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleTeacherFlow(telegramID, chatID int64, text string) {
	state := b.userStates[telegramID]

	if state.Step == "waiting_topic" {
		state.Topic = text
		state.Step = "waiting_count"

		msg := tgbotapi.NewMessage(chatID, "🔢 Nechta savol kerak? (1 dan 20 tagacha)")
		msg.ReplyMarkup = b.getBackMarkup()
		b.api.Send(msg)
		return
	}

	if state.Step == "waiting_count" {
		count, err := strconv.Atoi(text)
		if err != nil {
			b.api.Send(tgbotapi.NewMessage(chatID, "❌ Iltimos, faqat son kiriting (1-5)"))
			return
		}

		if count < 1 || count > 5 {
			b.api.Send(tgbotapi.NewMessage(chatID, "❌ Iltimos, 1 dan 5 gacha son kiriting"))
			return
		}

		// Arena queue ga yuborish
		requestData := map[string]interface{}{
			"telegram_id": telegramID,
			"prompt":       state.Topic,
			"count":       count,
		}

		err = rabbitmq.PublishArenaQueue(b.rabbitCh, requestData)
		if err != nil {
			log.Println("Arena queue ga yuborishda xato:", err)
			b.api.Send(tgbotapi.NewMessage(chatID, "❌ So‘rov yuborishda xato yuz berdi. Qayta urinib ko‘ring."))
			return
		}

		b.api.Send(tgbotapi.NewMessage(chatID, "✅ So‘rov yuborildi. Savollar tayyor bo‘lgach sizga xabar beriladi."))
		delete(b.userStates, telegramID)
		b.sendRoleMenu(chatID, "teacher")
		return
	}

	if state.Step == "waiting_test_id" {
		testData, err := b.resultService.GetUserTestResult(telegramID, text)
		if err != nil {
			b.api.Send(tgbotapi.NewMessage(chatID, "❌ Test topilmadi yoki hali tayyor emas"))
			return
		}

		msg := tgbotapi.NewMessage(chatID, b.formatTestResult(testData))
		msg.ParseMode = "HTML"
		b.api.Send(msg)

		delete(b.userStates, telegramID)
		b.sendRoleMenu(chatID, "teacher")
		return
	}
}