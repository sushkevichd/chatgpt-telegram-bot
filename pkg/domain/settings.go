package domain

const (
	ModelKey = "model"

	DefaultModel           = "gpt-4o-mini"
	SpeechRecognitionModel = "whisper-1"

	WelcomeMessage = `👋 Я твой ChatGPT Telegram-бот. Вот что умею:

❓ Отвечаю на вопросы. Напиши "новый чат" для очистки истории.
🎨 Рисую картинки. Начни запрос с "нарисуй".
🎙 Понимаю голосовые сообщения.
📷 Распознаю картинки.`
)

var SupportedModels = []string{
	"gpt-4o-mini",
	"gpt-4o",
	"gpt-4",
	"gpt-4-turbo",
	"gpt-3.5-turbo",
}

var DrawKeywords = []string{"рисуй", "draw"}
