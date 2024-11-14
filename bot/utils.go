package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(PrintEmoji, PrintEmoji+" Print All Tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(AddEmoji, " Add Task"),
			tgbotapi.NewInlineKeyboardButtonData(DeleteEmoji, " Delete Task"),
		),
	)
}
