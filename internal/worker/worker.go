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

		// Пропускаем, если это не команда
		if !job.Update.Message.IsCommand() {
			continue
		}

		cmd := job.Update.Message.Command()

		switch cmd {
		case "start":
			// Кнопка профиль (открывает мини-апп ВНУТРИ Телеграма)
			profileBtn := tgbotapi.NewInlineKeyboardButtonWebApp(
				"👤 Профиль",
				"https://ТВОЙ_МИНИ_АПП", // ← Вставь свою реальную ссылку!
			)

			// Остальные кнопки
			modelsBtn := tgbotapi.NewInlineKeyboardButtonData("🤖 GPT/Claude/Gemini", "models")
			designBtn := tgbotapi.NewInlineKeyboardButtonData("🎨 Дизайн с ИИ", "design_ai")
			audioBtn := tgbotapi.NewInlineKeyboardButtonData("🎵 Аудио с ИИ", "audio_ai")

			// Собираем клавиатуру
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(profileBtn),
				tgbotapi.NewInlineKeyboardRow(modelsBtn),
				tgbotapi.NewInlineKeyboardRow(designBtn, audioBtn),
			)

			// Создаём сообщение
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🤖 Добро пожаловать в CyberMate!\nВыберите раздел:")

			// Прикрепляем клавиатуру
			msg.ReplyMarkup = keyboard

			// Отправляем
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println("Ошибка отправки меню:", err)
			}

		case "help":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "📚 Доступные команды:\n• /start — главное меню\n• /models — список моделей")
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println(err)
			}

		case "models":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "🤖 Доступные модели:\n• GPT-4\n• Claude 3\n• Gemini Pro")
			if _, err := job.Bot.Send(msg); err != nil {
				log.Println(err)
			}

		case "info":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "ℹ️ CyberMate — твой персональный ИИ-помощник")
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
