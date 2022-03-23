package main

import (
	"log"

	"github.com/rikster17/TelegramBot1/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5161306872:AAFR2360W3NpwUnaa3IBrNiEjUfmT57BUH0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
