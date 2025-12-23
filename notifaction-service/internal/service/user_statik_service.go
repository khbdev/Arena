package service

import (
	"fmt"
	"log"

	"notifaction-service/internal/config"
	"notifaction-service/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendTestResult(result model.TestResult) {
	bot := config.TelegramBot()

	// =========================
	// 1️⃣ O‘QUVCHI UCHUN: Testlar soni va jami ball
	// =========================
	userResultText := fmt.Sprintf(
		"🏆 <b>Test natijasi</b>\n\n"+
			"📚 <b>Testlar soni:</b> %d\n"+
			"📊 <b>Jami ball:</b> %d\n\n"+
			"👏 <b>Omad!</b>",
		result.Total,
		result.Score,
	)

	userMsg := tgbotapi.NewMessage(result.TelegramID, userResultText)
	userMsg.ParseMode = "HTML"

	if _, err := bot.Send(userMsg); err != nil {
		log.Printf("❌ User result send error (%d): %v", result.TelegramID, err)
	} else {
		log.Printf("✅ Result sent to user %d", result.TelegramID)
	}

	// =========================
	// 2️⃣ O‘QITUVCHI UCHUN: Test ID, o‘quvchi ismi va ball
	// =========================
	teacherResultText := fmt.Sprintf(
		"📊 <b>Test natijasi</b>\n\n"+
			"🆔 <b>Test ID:</b> <code>%s</code>\n"+
			"👤 <b>O‘quvchi:</b> %s\n"+
			"✅ <b>Ball:</b> %d",
		result.TestID,
		result.UserFirstName,
		result.Score,
	)

	teacherMsg := tgbotapi.NewMessage(result.TeacherTelegramID, teacherResultText)
	teacherMsg.ParseMode = "HTML"

	if _, err := bot.Send(teacherMsg); err != nil {
		log.Printf("❌ Teacher result send error (%d): %v", result.TeacherTelegramID, err)
	} else {
		log.Printf("✅ Result sent to teacher %d", result.TeacherTelegramID)
	}
}
