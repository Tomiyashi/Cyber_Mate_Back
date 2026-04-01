package worker

import (
	"log"
	"time"

	"CyberMate_Back/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Worker обрабатывает задачи из канала
func Worker(ch <-chan models.Job) {
	for job := range ch {
		// Пропускаем, если нет сообщения
		if job.Update.Message == nil {
			continue
		}

		text := job.Update.Message.Text

		switch text {
		case "/start":
			btn1 := tgbotapi.NewKeyboardButton("👤 Профиль")
			btn1.WebApp = &tgbotapi.WebAppInfo{
				URL: "https://твой-mini-app.com",
			}
			btn2 := tgbotapi.NewKeyboardButton("🤖 GPT/Claude/Gemini")
			btn3 := tgbotapi.NewKeyboardButton("🎨 Дизайн с ИИ")
			btn4 := tgbotapi.NewKeyboardButton("🎵 Аудио с ИИ")
			btn5 := tgbotapi.NewKeyboardButton("📚 База знаний")

			row1 := []tgbotapi.KeyboardButton{btn1}
			row2 := []tgbotapi.KeyboardButton{btn2}
			row3 := []tgbotapi.KeyboardButton{btn3, btn4}
			row4 := []tgbotapi.KeyboardButton{btn5}

			keyboard := tgbotapi.ReplyKeyboardMarkup{
				Keyboard:       [][]tgbotapi.KeyboardButton{row1, row2, row3, row4},
				ResizeKeyboard: true,
			}

			// Создаём сообщение
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🤖 Добро пожаловать в CyberMate!\nВыберите раздел:")

			// Прикрепляем клавиатуру
			msg.ReplyMarkup = keyboard

			// Отправляем
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println("Ошибка отправки меню:", err)
			}

		case "👤 Профиль":
			// Отправляем временное сообщение
			tempMsg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🚀 Профиль открывается...")
			tempMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // Убираем клавиатуру

			sentMsg, err := job.Bot.Send(tempMsg)
			if err != nil {
				log.Println("Ошибка отправки сообщения:", err)
				return
			}

			// Запускаем горутину для удаления сообщения через 2 секунды
			go func(chatID int64, messageID int) {
				time.Sleep(2 * time.Second)
				_, err := job.Bot.Request(tgbotapi.DeleteMessage{
					ChatID:    chatID,
					MessageID: messageID,
				})
				if err != nil {
					log.Println("Ошибка удаления сообщения:", err)
				}
			}(job.Update.Message.Chat.ID, sentMsg.MessageID)

		case "🤖 GPT/Claude/Gemini":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🤖 Доступные модели:\n• GPT-4\n• Claude 3\n• Gemini Pro")
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println(err)
			}
		case "🎨 Дизайн с ИИ":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🎨 Дизайн с ИИ — в разработке 🚧")
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println(err)
			}
		case "🎵 Аудио с ИИ":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🎵 Аудио с ИИ — в разработке 🚧")
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println(err)
			}
		case "📚 База знаний":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "📚 База знаний:\n• /start — главное меню\n• /help — справка")
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println(err)
			}

		default:
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "❓ Неизвестная команда. Напиши /help для справки.")
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println(err)
			}
		}
	}
}
