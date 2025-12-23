// internal/bot/commands.go
package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCommand(update tgbotapi.Update) {
	cmd := update.Message.Command()
	switch cmd {
	case "start":
		b.handleStart(update)
	case "delete":
		b.handleDelete(update)
	}
}
func (b *Bot) handleDelete(update tgbotapi.Update) {
	telegramID := int64(update.Message.From.ID)
	chatID := update.Message.Chat.ID

	msgStr, err := b.telegramClient.CheckTelegramID(telegramID)
	if err != nil {
		b.api.Send(tgbotapi.NewMessage(chatID, "Xatolik yuz berdi"))
		return
	}

	b.api.Send(tgbotapi.NewMessage(chatID, msgStr))
}



func (b *Bot) handleStart(update tgbotapi.Update) {
	telegramID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	user, err := b.userService.GetByTelegramID(int64(telegramID))
	if err != nil {
		log.Println("GetByTelegramID error:", err)
		return
	}

	if user != nil {
		b.sendRoleMenu(chatID, user.Role)
		return
	}

	b.askForRole(chatID)
}

func (b *Bot) askForRole(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Siz Teacher yoki Student ekansiz?")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("👨‍🏫 Teacher", "role_teacher"),
			tgbotapi.NewInlineKeyboardButtonData("👨‍🎓 Student", "role_student"),
		),
	)
	b.api.Send(msg)
}

func (b *Bot) sendRoleMenu(chatID int64, role string) {
	var text string
	var rows [][]tgbotapi.InlineKeyboardButton

	switch role {
	case "teacher":
		text = "👨‍🏫 Teacher menyusi:"
		rows = [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("➕ Savol yaratish", "create_question"),
				tgbotapi.NewInlineKeyboardButtonData("📋 Savollarni ko‘rish", "view_questions"),
			},
		}
	case "student":
		text = "👨‍🎓 Student menyusi:"
		rows = [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("✅ Test yechish", "start_test"),
			},
		}
	default:
		text = "Noma’lum rol"
	}

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	b.api.Send(msg)
}

func (b *Bot) handleCallback(update tgbotapi.Update) {
	cb := update.CallbackQuery
	data := cb.Data
	chatID := cb.Message.Chat.ID
	telegramID := cb.From.ID

	// Loading o‘chirish
	_, err := b.api.Request(tgbotapi.NewCallback(cb.ID, ""))
	if err != nil {
		log.Printf("Callback acknowledge error: %v", err)
	}

	if data == "role_teacher" || data == "role_student" {
		role := "student"
		if data == "role_teacher" {
			role = "teacher"
		}

		_, err := b.userService.CreateUser(int64(telegramID), role, cb.From.FirstName)
		if err != nil {
			b.api.Send(tgbotapi.NewMessage(chatID, "Xato yuz berdi"))
			return
		}

		b.sendRoleMenu(chatID, role)
		return
	}

	if data == "create_question" {
		b.userStates[telegramID] = &UserState{Role: "teacher", Step: "waiting_topic"}
		msg := tgbotapi.NewMessage(chatID, "📝 Mavzuni kiriting:")
		msg.ReplyMarkup = b.getBackMarkup()
		b.api.Send(msg)
		return
	}

	if data == "view_questions" {
		b.userStates[telegramID] = &UserState{Role: "teacher", Step: "waiting_test_id"}
		msg := tgbotapi.NewMessage(chatID, "🆔 Test ID kiriting:")
		msg.ReplyMarkup = b.getBackMarkup()
		b.api.Send(msg)
		return
	}

	if data == "start_test" {
		b.userStates[telegramID] = &UserState{
			Role:      "student",
			Step:      "waiting_test_id",
			FirstName: cb.From.FirstName,
		}
		msg := tgbotapi.NewMessage(chatID, "🆔 Test ID kiriting:")
		msg.ReplyMarkup = b.getBackMarkup()
		b.api.Send(msg)
		return
	}

	if data == "back_to_menu" {
		user, err := b.userService.GetByTelegramID(telegramID)
		if err != nil || user == nil {
			b.api.Send(tgbotapi.NewMessage(chatID, "Xato yuz berdi"))
			return
		}
		delete(b.userStates, telegramID)
		b.sendRoleMenu(chatID, user.Role)
		return
	}
}