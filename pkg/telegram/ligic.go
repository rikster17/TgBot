package telegram

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart            = "start"
	commandHelp             = "help"
	commandTime             = "time"
	commandMyName           = "myname"
	commandFindTrackSpotify = "fmusic"
	spotifyURL              = "https://open.spotify.com/"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		if update.Message.IsCommand() {
			b.HandleCommand(update.Message)
			continue
		}
		b.bot.Send(msg)
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
		b.HandleCommandStart(message)
	case commandHelp:
		b.HandleCommandHelp(message)
	case commandTime:
		b.HandleTime(message)
	case commandMyName:
		b.HandleYourName(message)
	case commandFindTrackSpotify:
		b.HandleFindTrack(message)
	default:
		b.bot.Send(msg)
	}
}

func (b *Bot) HandleCommandStart(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome this is my first bot")
	b.bot.Send(msg)
}

func (b *Bot) HandleCommandHelp(message *tgbotapi.Message) {
	startInfo := fmt.Sprintf("/%s - start ToDo_Bot\n", commandStart)
	timeInfo := fmt.Sprintf("/%s - shows you current time +6\n", commandTime)
	mynameInfo := fmt.Sprintf("/%s - shows your nickname\n", commandMyName)
	findTrackInfo := fmt.Sprintf("/%s - find track by name", commandFindTrackSpotify)
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(startInfo, timeInfo, mynameInfo, findTrackInfo))
	b.bot.Send(msg)
}

func (b *Bot) HandleTime(message *tgbotapi.Message) {
	dt := time.Now()

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Your time is = %s", dt.Format(time.RFC1123)))
	b.bot.Send(msg)
}

func (b *Bot) HandleYourName(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.From.String())

	b.bot.Send(msg)
}

func (b *Bot) HandleFindTrack(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("This is your track"))

	b.bot.Send(msg)
}
