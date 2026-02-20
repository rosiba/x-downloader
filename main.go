package main

import (
	"log"
	"os"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	key := os.Getenv("TELEGRAM_BOT_API_KEY")
	if key == "" {
		log.Panic("TELEGRAM_BOT_API_KEY environment variable not set")
	}

	webHookURL := os.Getenv("TELEGRAM_WEBHOOK_URL")
	if webHookURL == "" {
		log.Panic("TELEGRAM_WEBHOOK_URL environment variable not set")
	}

	bot, err := botapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	wh, err := botapi.NewWebhook(webHookURL + "/" + bot.Token)
	if err != nil {
		log.Panic(err)
	}

	if _, err := bot.Request(wh); err != nil {
		log.Panic(err)
	}

	log.Println("Bot is now running.")
}
