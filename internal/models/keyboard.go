package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetMainKeyboard(webAppURL string) tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("👤 Профиль")
	btn1.WebApp = &tgbotapi.WebAppInfo{URL: webAppURL}

	btn2 := tgbotapi.NewKeyboardButton("⚙️ Настройки")
	btn3 := tgbotapi.NewKeyboardButton("🛟 Помощь")

	row1 := []tgbotapi.KeyboardButton{btn1}
	row2 := []tgbotapi.KeyboardButton{btn2, btn3}

	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:       [][]tgbotapi.KeyboardButton{row1, row2},
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

func GetSettingsKeyboard() *tgbotapi.InlineKeyboardMarkup {
	btnLang := tgbotapi.NewInlineKeyboardButtonData("🌐 Язык", "Settings_Lang")
	btnBack := tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", "Main_Back")

	row1 := []tgbotapi.InlineKeyboardButton{btnLang}
	row2 := []tgbotapi.InlineKeyboardButton{btnBack}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)
	return &keyboard
}

func GetLanguageKeyboard() *tgbotapi.InlineKeyboardMarkup {
	btnRu := tgbotapi.NewInlineKeyboardButtonData("🇷🇺 Русский", "Lang_ru")
	btnEn := tgbotapi.NewInlineKeyboardButtonData("🇬🇧 English", "Lang_en")
	btnBack := tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", "Settings_Back")

	row1 := []tgbotapi.InlineKeyboardButton{btnRu, btnEn}
	row2 := []tgbotapi.InlineKeyboardButton{btnBack}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)
	return &keyboard
}
