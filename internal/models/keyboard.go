package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// GetMainKeyboard возвращает главное меню
func GetMainKeyboard(webAppURL string) tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("👤 Профиль")
	btn1.WebApp = &tgbotapi.WebAppInfo{URL: webAppURL}

	btn2 := tgbotapi.NewKeyboardButton("🤖 Нейросети")
	btn3 := tgbotapi.NewKeyboardButton("🛟 Помощь")

	row1 := []tgbotapi.KeyboardButton{btn1}
	row2 := []tgbotapi.KeyboardButton{btn2, btn3}

	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:       [][]tgbotapi.KeyboardButton{row1, row2},
		ResizeKeyboard: true,
	}
}

// GetNeuroKeyboard возвращает меню нейросетей
func GetNeuroKeyboard() tgbotapi.ReplyKeyboardMarkup {
	ai1 := tgbotapi.NewKeyboardButton("🎨 Дизайн с ИИ")
	ai2 := tgbotapi.NewKeyboardButton("🎵 Аудио с ИИ")
	ai3 := tgbotapi.NewKeyboardButton("🤖 GPT/Claude/Gemini")
	ai4 := tgbotapi.NewKeyboardButton("⬅️ Назад")

	row1 := []tgbotapi.KeyboardButton{ai1, ai2}
	row2 := []tgbotapi.KeyboardButton{ai3}
	row3 := []tgbotapi.KeyboardButton{ai4}

	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:       [][]tgbotapi.KeyboardButton{row1, row2, row3},
		ResizeKeyboard: true,
	}
}

func GetSupportKeyboard() tgbotapi.ReplyKeyboardMarkup {
	sup1 := tgbotapi.NewKeyboardButton("💬 Написать в поддержку")
	sup2 := tgbotapi.NewKeyboardButton("⬅️ Назад")

	row1 := []tgbotapi.KeyboardButton{sup1}
	row2 := []tgbotapi.KeyboardButton{sup2}

	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:       [][]tgbotapi.KeyboardButton{row1, row2},
		ResizeKeyboard: true,
	}
}
