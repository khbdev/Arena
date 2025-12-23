// internal/bot/student.go
package bot

import (
	"geteway/internal/model"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleStudentFlow(telegramID, chatID int64, text string) {
	state := b.userStates[telegramID]

	if state.Step == "waiting_test_id" {
		testData, err := b.resultService.GetUserTestResult(telegramID, text)
		if err != nil {
			b.api.Send(tgbotapi.NewMessage(chatID, "❌ Test topilmadi yoki mavjud emas"))
			return
		}

		state.Step = "solving_test"
		state.TestID = text
		state.Questions = testData.Questions
		state.Current = 0
		state.Answers = []model.UserAnswer{}

		b.sendQuestion(chatID, telegramID)
		return
	}

	if state.Step == "solving_test" {
		text = strings.ToUpper(strings.TrimSpace(text))
		if text == "A" || text == "B" || text == "C" || text == "D" {
			q := state.Questions[state.Current]

			state.Answers = append(state.Answers, model.UserAnswer{
				QuestionID:    q.ID,
				CorrectAnswer: text,
			})

			state.Current++

			if state.Current >= len(state.Questions) {
				b.finishTest(chatID, telegramID)
				return
			}

			b.sendQuestion(chatID, telegramID)
			return
		}

		b.api.Send(tgbotapi.NewMessage(chatID, "⚠️ Faqat A, B, C yoki D variantini tanlang!"))
	}
}