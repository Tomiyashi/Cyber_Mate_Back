package handler

import (
	"CyberMate_Back/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot        *tgbotapi.BotAPI
	jobs       chan<- models.Job
	userRepo   models.UserRepository
	miniAppURL string
}

func New(bot *tgbotapi.BotAPI, jobs chan<- models.Job, userRepo models.UserRepository, miniAppURL string) *Handler {
	return &Handler{
		bot:        bot,
		jobs:       jobs,
		userRepo:   userRepo,
		miniAppURL: miniAppURL,
	}
}

func (h *Handler) Handle(update tgbotapi.Update) {
	if update.Message == nil && update.CallbackQuery == nil {
		return
	}

	h.jobs <- models.Job{
		Bot:        h.bot,
		Update:     update,
		UserRepo:   h.userRepo,
		MiniAppURL: h.miniAppURL,
	}
}
