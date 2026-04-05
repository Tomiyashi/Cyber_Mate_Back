package worker

import (
	"log"

	"CyberMate_Back/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Worker(ch <-chan models.Job) {
	for job := range ch {
		if job.Update.Message == nil {
			continue
		}

		// Объявляем переменные здесь, чтобы они были видны во всем цикле
		chatID := job.Update.Message.Chat.ID
		text := job.Update.Message.Text
		switch text {
		case "/start":
			btn1 := tgbotapi.NewKeyboardButton("👤 Профиль")
			btn1.WebApp = &tgbotapi.WebAppInfo{
				URL: "https://твой-mini-app.com",
			}
			btn2 := tgbotapi.NewKeyboardButton("🤖 Нейросети")
			btn3 := tgbotapi.NewKeyboardButton("🛟 Помощь")

			row1 := []tgbotapi.KeyboardButton{btn1}
			row2 := []tgbotapi.KeyboardButton{btn2, btn3}

			keyboard := tgbotapi.ReplyKeyboardMarkup{
				Keyboard:       [][]tgbotapi.KeyboardButton{row1, row2},
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
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🚀 Профиль открывается...")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println("Ошибка отправки подтверждения профиля:", err)
			}

		case "🤖 Нейросети":
			msg := tgbotapi.NewMessage(chatID, "🧠 Выбери нейросеть:")
			msg.ReplyMarkup = models.GetNeuroKeyboard()
			job.Bot.Send(msg)

		case "🛟 Помощь":
			msg := tgbotapi.NewMessage(chatID, "🛟 Нужна помощь?")
			msg.ReplyMarkup = models.GetSupportKeyboard()
			job.Bot.Send(msg)

		case "⬅️ Назад":
			msg := tgbotapi.NewMessage(chatID, "🏠 Главное меню")
			msg.ReplyMarkup = models.GetMainKeyboard("https://твой-mini-app.com")
			job.Bot.Send(msg)

		default:
			msg := tgbotapi.NewMessage(chatID, "❓ Я тебя не понял.")
			job.Bot.Send(msg)
		}
	}
}
