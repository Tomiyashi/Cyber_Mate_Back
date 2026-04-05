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
		// Проверяем, что сообщение вообще есть
		if update.Message == nil {
			continue
		}

		log.Println("Получено сообщение:", update.Message.Text)

		// МЫ УБРАЛИ IsCommand(). Теперь любое сообщение (и /start, и текст кнопок)
		// отправляется в канал воркерам на обработку.
		ch <- models.Job{
			Bot:    bot,
			Update: update,
		}
	}
}
