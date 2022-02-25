package main

import (
	"github.com/rikster17/TelegramBot1/pkg/telegram"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5288554328:AAEOcSKTYyeROthTNqCoT5UXcpnLYAuX8vQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
