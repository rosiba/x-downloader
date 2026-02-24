package main

import (
	"log"
	"net/http"
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

	port := os.Getenv("PORT")
	if port == "" {
		log.Panic("PORT environment variable not set")
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

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Panic(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:"+port, nil)
	log.Println("Bot is now running.")

	for update := range updates {
		log.Printf("%+v\n", update)
		if update.Message == nil {
			url, err := getURL(update.Message.Text)
			if err != nil {
				msg := botapi.NewMessage(update.Message.Chat.ID, "Failed to fetch link")
				bot.Send(msg)
				continue

			}
			msg := botapi.NewMessage(update.Message.Chat.ID, url)
			bot.Send(msg)
		}
	}
}
