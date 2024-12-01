package main

import (
    "log"
    "os"
    
    "telegram-bot/bot"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
    token := os.Getenv("TELEGRAM_TOKEN")
    if token == "" {
        log.Fatal("TELEGRAM_TOKEN environment variable is not set")
    }

    botAPI, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        log.Fatal(err)
    }

    handler := bot.NewConversationHandler(botAPI)
    handler.Start()
} 