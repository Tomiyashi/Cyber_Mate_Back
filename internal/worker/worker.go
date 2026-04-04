package worker

import (
	"log"

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
			keyboard := models.GetNeuroKeyboard()

			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🧠 Выбери нейросеть:")
			msg.ReplyMarkup = keyboard

			if _, err := job.Bot.Send(msg); err != nil {
				log.Println("Ошибка отправки меню нейросетей:", err)
			}
		case "🛟 Помощь":
			// Отправляем сообщение с поддержкой
			msg := tgbotapi.NewMessage(
				job.Update.Message.Chat.ID,
				"🛟 Служба поддержки:\nНапиши нам: @CyberMate_Support",
			)
			msg.ReplyMarkup = models.GetSupportKeyboard()

			if _, err := job.Bot.Send(msg); err != nil {
				log.Println("Ошибка отправки сообщения поддержки:", err)
			}

		case "🎨 Дизайн с ИИ":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🎨 Дизайн с ИИ — в разработке 🚧")
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println(err)
			}

		case "💬 Написать в поддержку":
			msg := tgbotapi.NewMessage(
				job.Update.Message.Chat.ID,
				"📩 Напишите нам: @CyberMate_Support",
			)

			btn := tgbotapi.NewInlineKeyboardButtonURL("✈️ Открыть чат", "https://t.me/CyberMate_Support")
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(btn))
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println("Ошибка отправки ссылки:", err)
			}

		case "⬅️ Назад":
			keyboard := models.GetMainKeyboard("https://твой-mini-app.com")
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🤖 Добро пожаловать в CyberMate!\nВыберите раздел:")
			msg.ReplyMarkup = keyboard
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println("Ошибка отправки меню:", err)
			}

		default:
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "❓ Неизвестная команда.")
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println(err)
			}
		}
	}
}
