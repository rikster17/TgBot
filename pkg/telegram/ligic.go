package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	commandStart = "start"
	commandHelp  = "help"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		b.bot.Send(msg)

		if update.Message.IsCommand() {
			b.HandleCommand(update.Message)
			continue
		}
	}
}

func (b *Bot) InitUpdatesChannel() tgbotapi.UpdatesChannel {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) HandleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know this command :(\nYou can see other commands with /help")

	switch message.Command() {
	case commandStart:
		msg.Text = "Welcome this is my first bot"
		b.bot.Send(msg)
	case commandHelp:
		msg.Text = "/start - start ToDo_Bot"
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
