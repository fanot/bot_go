//nolint:unused
package bot

import (
	"fmt"
	"strconv"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ConversationHandler struct {
	bot       *tgbotapi.BotAPI
	userState map[int64]State
	userData  map[int64]*UserData
	mu        sync.RWMutex
}

func NewConversationHandler(bot *tgbotapi.BotAPI) *ConversationHandler {
	return &ConversationHandler{
		bot:       bot,
		userState: make(map[int64]State),
		userData:  make(map[int64]*UserData),
	}
}

func (h *ConversationHandler) handleStart(message *tgbotapi.Message) {
	h.mu.Lock()
	h.userState[message.From.ID] = StateChoosing
	if _, exists := h.userData[message.From.ID]; !exists {
		h.userData[message.From.ID] = &UserData{
			Facts: make([]string, 0),
		}
	}
	h.mu.Unlock()

	msg := tgbotapi.NewMessage(message.Chat.ID, ChooseOptionMessage)
	h.bot.Send(msg)
}

func (h *ConversationHandler) HandleMessage(message *tgbotapi.Message) {
	fmt.Printf("Received message from user %d: %s\n", message.From.ID, message.Text)

	h.mu.RLock()
	state, exists := h.userState[message.From.ID]
	h.mu.RUnlock()

	fmt.Printf("User state: %v, exists: %v\n", state, exists)

	if !exists {
		h.handleStart(message)
		return
	}

	switch state {
	case StateChoosing:
		h.handleChoice(message)
	case StateTypingReply:
		h.handleReply(message)
	}
}

func (h *ConversationHandler) handleChoice(message *tgbotapi.Message) {
	choice, err := strconv.Atoi(message.Text)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, InvalidOptionMessage)
		h.bot.Send(msg)
		return
	}

	var replyText string
	switch choice {
	case 1:
		replyText = AskForFactMessage
	case 2:
		replyText = AskForLocationMessage
	case 3:
		replyText = AskForBioMessage
	case 4:
		h.handleShowData(message)
		return
	default:
		replyText = InvalidOptionMessage
		h.bot.Send(tgbotapi.NewMessage(message.Chat.ID, replyText))
		return
	}

	h.mu.Lock()
	h.userState[message.From.ID] = StateTypingReply
	h.mu.Unlock()

	msg := tgbotapi.NewMessage(message.Chat.ID, replyText)
	h.bot.Send(msg)
}

func (h *ConversationHandler) handleReply(message *tgbotapi.Message) {
	h.mu.Lock()
	userData := h.userData[message.From.ID]
	if userData == nil {
		userData = &UserData{Facts: make([]string, 0)}
		h.userData[message.From.ID] = userData
	}

	// Сохраняем ответ пользователя
	userData.Facts = append(userData.Facts, message.Text)
	h.userState[message.From.ID] = StateChoosing
	h.mu.Unlock()

	// Отправляем подтверждение и показываем меню снова
	msg := tgbotapi.NewMessage(message.Chat.ID, DataSavedMessage+"\n\n"+ChooseOptionMessage)
	h.bot.Send(msg)
}

func (h *ConversationHandler) handleShowData(message *tgbotapi.Message) {
	h.mu.RLock()
	userData := h.userData[message.From.ID]
	h.mu.RUnlock()

	if userData == nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "У вас пока нет сохраненных данных")
		h.bot.Send(msg)
		return
	}

	text := "Ваши данные:\n\nФакты:\n"
	for i, fact := range userData.Facts {
		text += fmt.Sprintf("%d. %s\n", i+1, fact)
	}

	if userData.Location != "" {
		text += "\nМестоположение: " + userData.Location
	}
	if userData.Bio != "" {
		text += "\nБио: " + userData.Bio
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	h.bot.Send(msg)
}

func (h *ConversationHandler) Start() {
	fmt.Println("ConversationHandler started")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.bot.GetUpdatesChan(u)

	// Бесконечный цикл обработки сообщений
	for update := range updates {
		if update.Message == nil {
			continue
		}
		h.HandleMessage(update.Message)
	}
}

func (h *ConversationHandler) GetUserState(userID int64) State {
	if state, exists := h.userState[userID]; exists {
		return state
	}
	return StateChoosing // Default state
}
