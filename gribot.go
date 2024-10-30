package main

import (
	"log"
	"net/http"
	"os"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Входная точка приложения
func main() {
	// Получаем токен бота из переменных окружения
	botToken := os.Getenv("7647349322:AAHT6rp0Jz1e1QwacZYMbtIjiAlGNQxV2VY")
	if botToken == "" {
		log.Fatal("Токен бота не установлен")
	}

	// Создаем новый бот
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	// Логируем, что бот запущен
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настраиваем обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Получаем обновления
	updates, err := bot.GetUpdatesChan(u)

	// Обрабатываем обновления
	for update := range updates {
		if update.Message == nil { // игнорируем сообщения без текста
			continue
		}

		// Логируем входящее сообщение
		log.Printf("Received message: %s", update.Message.Text)

		// Если пользователь отправил команду /start, отвечаем приветствием
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я GriBot: Мрачный шутник. Как я могу помочь?")
				bot.Send(msg)
			}
		}

		// Добавляем простую функцию, которая отвечает на любое текстовое сообщение
		replyMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы сказали: "+update.Message.Text)
		bot.Send(replyMsg)

		// Дополнительная простая функция, которая отвечает случайной шуткой
		if update.Message.Text == "шутка" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Почему программисты предпочитают темный фон? Потому что свет притягивает жуков!")
			bot.Send(msg)
		}
	}
}

