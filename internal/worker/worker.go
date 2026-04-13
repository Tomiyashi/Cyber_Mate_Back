package worker

import (
	"context"
	"log"
	"time"

	"CyberMate_Back/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func t(lang, ru, en string) string {
	if lang == "en" {
		return en
	}
	return ru
}

const dbTimeout = 3 * time.Second

func withDBTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), dbTimeout)
}

func Worker(ch <-chan models.Job) {
	for job := range ch {
		chatID := int64(0)

		if job.Update.CallbackQuery != nil {
			if job.Update.CallbackQuery.Message != nil {
				chatID = job.Update.CallbackQuery.Message.Chat.ID
			}
		} else if job.Update.Message != nil {
			chatID = job.Update.Message.Chat.ID
		}

		if chatID != 0 {
			ctx, cancel := withDBTimeout()
			err := job.UserRepo.Upsert(ctx, chatID)
			cancel()
			if err != nil {
				log.Printf("⚠️  Failed to upsert user %d: %v", chatID, err)
			}
		}

		if job.Update.CallbackQuery != nil {
			cb := job.Update.CallbackQuery
			if cb.Message == nil {
				if _, err := job.Bot.Send(tgbotapi.NewCallback(cb.ID, "")); err != nil {
					log.Printf("❌ Callback error: %v", err)
				}
				continue
			}

			chatID := cb.Message.Chat.ID
			messageID := cb.Message.MessageID
			callbackID := cb.ID
			data := cb.Data

			edit := tgbotapi.NewEditMessageText(chatID, messageID, "")
			lang := "ru"
			ctx, cancel := withDBTimeout()
			user, err := job.UserRepo.Get(ctx, chatID)
			cancel()
			if err != nil {
				log.Printf("⚠️  Failed to get user settings: %v", err)
			} else if user != nil && user.Language != "" {
				lang = user.Language
			}

			switch data {
			case "Settings_Lang":
				edit.Text = "🌐 Выберите язык:"
				edit.ReplyMarkup = models.GetLanguageKeyboard()

			case "Settings_Back", "Lang_back":
				edit.Text = "⚙️ Настройки бота:"
				edit.ReplyMarkup = models.GetSettingsKeyboard()

			case "Main_Back":
				msg := tgbotapi.NewMessage(chatID, t(lang, "🤖 Добро пожаловать в CyberMate!\nВыберите раздел:", "🤖 Welcome to CyberMate!\nChoose a section:"))
				msg.ReplyMarkup = models.GetMainKeyboard(job.MiniAppURL)
				if _, err := job.Bot.Send(msg); err != nil {
					log.Printf("❌ Send message error: %v", err)
				}
				if _, err := job.Bot.Send(tgbotapi.NewCallback(callbackID, "")); err != nil {
					log.Printf("❌ Callback error: %v", err)
				}
				continue

			case "Lang_ru":
				ctx, cancel := withDBTimeout()
				err := job.UserRepo.UpdateLanguage(ctx, chatID, "ru")
				cancel()
				if err != nil {
					edit.Text = "❌ Ошибка при сохранении языка"
				} else {
					edit.Text = "✅ Язык изменён на Русский"
				}
				edit.ReplyMarkup = models.GetLanguageKeyboard()

			case "Lang_en":
				ctx, cancel := withDBTimeout()
				err := job.UserRepo.UpdateLanguage(ctx, chatID, "en")
				cancel()
				if err != nil {
					edit.Text = "❌ Error saving language"
				} else {
					edit.Text = "✅ Language changed to English"
				}
				edit.ReplyMarkup = models.GetLanguageKeyboard()

			default:
				edit.Text = "⚠️ Неизвестное действие"
			}

			if _, err := job.Bot.Send(edit); err != nil {
				log.Printf("❌ Edit message error: %v", err)
			}
			if _, err := job.Bot.Send(tgbotapi.NewCallback(callbackID, "")); err != nil {
				log.Printf("❌ Callback error: %v", err)
			}
			continue
		}

		if job.Update.Message == nil {
			continue
		}

		ctx, cancel := withDBTimeout()
		user, err := job.UserRepo.Get(ctx, job.Update.Message.Chat.ID)
		cancel()
		if err != nil {
			log.Printf("⚠️  Failed to get user settings: %v", err)
			user = &models.User{Language: "ru"}
		}

		lang := user.Language
		text := job.Update.Message.Text

		switch text {
		case "/start":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, t(lang, "🤖 Добро пожаловать в CyberMate!\nВыберите раздел:", "🤖 Welcome to CyberMate!\nChoose a section:"))
			msg.ReplyMarkup = models.GetMainKeyboard(job.MiniAppURL)
			if _, err := job.Bot.Send(msg); err != nil {
				log.Printf("❌ Send menu error: %v", err)
			}

		case "👤 Профиль":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, t(lang, "🚀 Профиль открывается...", "🚀 Opening profile..."))
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := job.Bot.Send(msg); err != nil {
				log.Printf("❌ Send profile error: %v", err)
			}

		case "⚙️ Настройки":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, "⚙️ Настройки бота:")
			msg.ReplyMarkup = models.GetSettingsKeyboard()
			if _, err := job.Bot.Send(msg); err != nil {
				log.Printf("❌ Send settings error: %v", err)
			}

		case "🛟 Помощь":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, t(lang, "🛟 Служба поддержки:\nНапиши нам: @CyberMate_Support", "🛟 Support:\nWrite to us: @CyberMate_Support"))
			msg.ReplyMarkup = models.GetSupportKeyboard()
			if _, err := job.Bot.Send(msg); err != nil {
				log.Printf("❌ Send support error: %v", err)
			}

		case "💬 Написать в поддержку":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, t(lang, "📩 Напишите нам: @CyberMate_Support", "📩 Write to us: @CyberMate_Support"))
			btn := tgbotapi.NewInlineKeyboardButtonURL("✈️ Открыть чат", "https://t.me/CyberMate_Support")
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(btn))
			if _, err := job.Bot.Send(msg); err != nil {
				log.Printf("❌ Send link error: %v", err)
			}

		case "⬅️ Назад":
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, t(lang, "🤖 Добро пожаловать в CyberMate!\nВыберите раздел:", "🤖 Welcome to CyberMate!\nChoose a section:"))
			msg.ReplyMarkup = models.GetMainKeyboard(job.MiniAppURL)
			if _, err := job.Bot.Send(msg); err != nil {
				log.Printf("❌ Send back menu error: %v", err)
			}

		default:
			msg := tgbotapi.NewMessage(job.Update.Message.Chat.ID, t(lang, "❓ Неизвестная команда.", "❓ Unknown command."))
			if _, err := job.Bot.Send(msg); err != nil {
				log.Printf("❌ Send unknown command error: %v", err)
			}
		}
	}
}
