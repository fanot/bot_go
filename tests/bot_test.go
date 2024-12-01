package tests

import (
	"testing"
	
	"telegram-bot/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TestConversationHandler_Start(t *testing.T) {
	mockBot := &tgbotapi.BotAPI{}
	handler := bot.NewConversationHandler(mockBot)

	// Test initial state
	if state := handler.GetUserState(123); state != bot.StateChoosing {
		t.Errorf("Expected initial state to be StateChoosing, got %v", state)
	}
}

