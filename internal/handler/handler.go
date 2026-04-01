package handler

import (
	"CyberMate_Back/internal/models"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(
	updates <-chan tgbotapi.Update,
	ch chan<- models.Job,
	bot *tgbotapi.BotAPI,
) {
	for update := range updates {
		if update.Message != nil {
			log.Println("Получено сообщение:", update.Message.Text)
		}
		if ok := update.Message.IsCommand(); ok {
			ch <- models.Job{
				Bot:    bot,
				Update: update,
			}
		}
	}
}
