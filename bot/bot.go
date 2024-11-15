package bot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type State int

const (
	StateStart State = iota
	StatePrint
	StateAdd
	StateDelete
)

type Bot struct {
	bot   *tgbotapi.BotAPI
	state State
}

func (b *Bot) Run() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		return
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := b.bot.GetUpdates(updateConfig)
	if err != nil {
		panic(err)
	}

	for _, update := range updates {
		if update.Message != nil {
			switch b.state {
			case StateStart:
				b.handleStart(&update)
				// case StatePrint:
				// 	b.handlePrint(update.Message, &update)
				// case StateAdd:
				// 	b.handleAdd(update.Message, &update)
				// case StateDelete:
				// 	b.handleDelete(update.Message, &update)
			}
		}
	}
}

func (b *Bot) handleStart(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if update.Message.Text == "/start" {
		msg.Text = StartMessage
		msg.ReplyMarkup = CreateInlineKeyboard()

		if msg.Text != "" {
			if _, err := b.bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
