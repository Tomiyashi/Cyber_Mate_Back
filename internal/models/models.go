package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Job — задача для воркера
type Job struct {
	Bot    *tgbotapi.BotAPI
	Update tgbotapi.Update
}
