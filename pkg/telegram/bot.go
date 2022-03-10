package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zmb3/spotify"
	"log"
)

type Bot struct {
	bot            *tgbotapi.BotAPI
	spotifyTracker *spotify.FullTrack
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.InitUpdatesChannel()
	b.handleUpdates(updates)
	return nil
}
