package models

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type User struct {
	ChatID   int64  `json:"chat_id"`
	Language string `json:"language"`
}

type Job struct {
	Bot        *tgbotapi.BotAPI
	Update     tgbotapi.Update
	UserRepo   UserRepository
	MiniAppURL string
}

type UserRepository interface {
	Upsert(ctx context.Context, chatID int64) error
	Get(ctx context.Context, chatID int64) (*User, error)
	UpdateLanguage(ctx context.Context, chatID int64, lang string) error
}
