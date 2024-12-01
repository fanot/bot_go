package bot

// UserData представляет собой структуру для хранения данных пользователя
type UserData struct {
	Facts    []string
	Location string
	Bio      string
}

// Constants для команд и сообщений
const (
	StartMessage = "Привет! Я бот для сохранения информации. Что бы вы хотели сделать?"

	ChooseOptionMessage = `Выберите опцию:
1. Добавить факт о себе
2. Указать местоположение
3. Добавить фио
4. Показать сохраненные данные`

	AskForFactMessage     = "Пожалуйста, поделитесь интересным фактом о себе"
	AskForLocationMessage = "Где вы находитесь?"
	AskForBioMessage      = "Расскажите немного о себе"

	DataSavedMessage     = "Информация сохранена!"
	InvalidOptionMessage = "Пожалуйста, выберите существующую опцию"
)

// Reply представляет собой структуру для ответа бота
type Reply struct {
	Text        string
	ReplyMarkup interface{}
}

// Option представляет собой вариант выбора в меню
type Option struct {
	ID    int
	Label string
}

// Menu содержит все доступные опции
var Menu = []Option{
	{ID: 1, Label: "Добавить факт"},
	{ID: 2, Label: "Указать местоположение"},
	{ID: 3, Label: "Добавить био"},
	{ID: 4, Label: "Показать данные"},
}

type State int

const (
	StateChoosing State = iota
	StateTypingReply
)
