// internal/bot/question.go
package bot

import (
	"fmt"
	"log"
	"strings"

	"geteway/internal/model"
	rabbitmq "geteway/internal/rabbitMq"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleText(update tgbotapi.Update) {
	telegramID := update.Message.From.ID
	chatID := update.Message.Chat.ID
	text := strings.TrimSpace(update.Message.Text)

	state, ok := b.userStates[telegramID]
	if !ok || state == nil {
		return
	}

	if state.Role == "teacher" {
		b.handleTeacherFlow(telegramID, chatID, text)
	} else if state.Role == "student" {
		b.handleStudentFlow(telegramID, chatID, text)
	}
}

func (b *Bot) sendQuestion(chatID, telegramID int64) {
	state := b.userStates[telegramID]
	if state == nil || state.Current >= len(state.Questions) {
		return
	}

	q := state.Questions[state.Current]

	questionText := fmt.Sprintf(
		"<b>%d/%d</b>\n\n<b>%s</b>\n\n",
		state.Current+1,
		len(state.Questions),
		q.Question,
	)

	var optionsText strings.Builder
	for i, opt := range q.Options {
		optionsText.WriteString(fmt.Sprintf("%c) %s\n", 'A'+i, opt.Text))
	}
	questionText += optionsText.String()

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("A"),
			tgbotapi.NewKeyboardButton("B"),
			tgbotapi.NewKeyboardButton("C"),
			tgbotapi.NewKeyboardButton("D"),
		),
	)
	keyboard.OneTimeKeyboard = true
	keyboard.ResizeKeyboard = true

	msg := tgbotapi.NewMessage(chatID, questionText)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}

func (b *Bot) finishTest(chatID, telegramID int64) {
	state := b.userStates[telegramID]
	if state == nil {
		return
	}

	payload := model.TestAnswerRequest{
		TelegramID:    telegramID,
		TestID:        state.TestID,
		Answers:       state.Answers,
		UserFirstName: state.FirstName,
	}

	err := rabbitmq.PublishResultQueue(b.rabbitCh, payload)
	if err != nil {
		log.Println("Natijani yuborishda xato:", err)
		b.api.Send(tgbotapi.NewMessage(chatID, "❌ Natijani yuborishda xato yuz berdi"))
		return
	}

	successMsg := tgbotapi.NewMessage(chatID, "✅ Test yakunlandi!\nJavoblaringiz qabul qilindi va baholanmoqda.")
	successMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	b.api.Send(successMsg)

	delete(b.userStates, telegramID)

	user, _ := b.userService.GetByTelegramID(telegramID)
	if user != nil {
		b.sendRoleMenu(chatID, user.Role)
	}
}