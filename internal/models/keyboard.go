package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetMainKeyboard(webAppURL, lang string) tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton(mapText(lang, "👤 Профиль", "👤 Profile"))
	btn1.WebApp = &tgbotapi.WebAppInfo{URL: webAppURL}

	btn2 := tgbotapi.NewKeyboardButton(mapText(lang, "⚙️ Настройки", "⚙️ Settings"))
	btn3 := tgbotapi.NewKeyboardButton(mapText(lang, "🛟 Помощь", "🛟 Help"))

	row1 := []tgbotapi.KeyboardButton{btn1}
	row2 := []tgbotapi.KeyboardButton{btn2, btn3}

	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:       [][]tgbotapi.KeyboardButton{row1, row2},
		ResizeKeyboard: true,
	}
}

func GetSupportKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	sup1 := tgbotapi.NewKeyboardButton(mapText(lang, "💬 Написать в поддержку", "💬 Contact support"))
	sup2 := tgbotapi.NewKeyboardButton(mapText(lang, "⬅️ Назад", "⬅️ Back"))

	row1 := []tgbotapi.KeyboardButton{sup1}
	row2 := []tgbotapi.KeyboardButton{sup2}

	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:       [][]tgbotapi.KeyboardButton{row1, row2},
		ResizeKeyboard: true,
	}
}

func GetSettingsKeyboard(lang string) *tgbotapi.InlineKeyboardMarkup {
	btnLang := tgbotapi.NewInlineKeyboardButtonData(mapText(lang, "🌐 Язык", "🌐 Language"), "Settings_Lang")
	btnBack := tgbotapi.NewInlineKeyboardButtonData(mapText(lang, "⬅️ Назад", "⬅️ Back"), "Main_Back")

	row1 := []tgbotapi.InlineKeyboardButton{btnLang}
	row2 := []tgbotapi.InlineKeyboardButton{btnBack}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)
	return &keyboard
}

func GetLanguageKeyboard(lang string) *tgbotapi.InlineKeyboardMarkup {
	btnRu := tgbotapi.NewInlineKeyboardButtonData("🇷🇺 Русский", "Lang_ru")
	btnEn := tgbotapi.NewInlineKeyboardButtonData("🇬🇧 English", "Lang_en")
	btnBack := tgbotapi.NewInlineKeyboardButtonData(mapText(lang, "⬅️ Назад", "⬅️ Back"), "Settings_Back")

	row1 := []tgbotapi.InlineKeyboardButton{btnRu, btnEn}
	row2 := []tgbotapi.InlineKeyboardButton{btnBack}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)
	return &keyboard
}

func mapText(lang, ru, en string) string {
	if lang == "en" {
		return en
	}
	return ru
}
