package bot

import (
	"fmt"
	"strings"

	"geteway/internal/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) getBackMarkup() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Orqaga", "back_to_menu"),
		),
	)
}

func (b *Bot) formatTestResult(data *model.TestData) string {
	var sb strings.Builder

	sb.WriteString("<b>📋 Test savollari va to‘g‘ri javoblari</b>\n\n")
	sb.WriteString(fmt.Sprintf(
		"<b>Test ID:</b> <code>%s</code>\n\n",
		data.TestID,
	))

	for i, q := range data.Questions {
		// Savol
		sb.WriteString(fmt.Sprintf(
			"<b>%d. %s</b>\n",
			i+1,
			q.Question,
		))

		// Variantlar
		for j, opt := range q.Options {
			prefix := ""

			// TO‘G‘RI YECHIM (string(int) YO‘Q)
			if len(q.Correct) > 0 && q.Correct[0] == byte('A'+j) {
				prefix = "✅ "
			}

			sb.WriteString(fmt.Sprintf(
				"%s%c) %s\n",
				prefix,
				'A'+j,
				opt.Text,
			))
		}

		sb.WriteString("\n")
	}

	return sb.String()
}
