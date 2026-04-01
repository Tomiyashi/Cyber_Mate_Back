package main

import (
	// ✅ Внутренние пакеты проекта (твоя архитектура)
	"CyberMate_Back/internal/handler" // Приём сообщений
	"CyberMate_Back/internal/models"  // Общие структуры
	"CyberMate_Back/internal/worker"  // Обработка задач

	// ✅ Стандартная библиотека
	"log"
	"os"

	// ✅ Внешние зависимости
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// ─────────────────────────────────────────────────────
	// 1. КОНФИГУРАЦИЯ
	// ─────────────────────────────────────────────────────
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Режим разработки (без сети)
	if os.Getenv("DEV_MODE") == "true" {
		log.Println("DEV MODE: Bot started without network connection")
		log.Println("Set DEV_MODE=false to connect to Telegram")
		return
	}

	// ─────────────────────────────────────────────────────
	// 2. ИНИЦИАЛИЗАЦИЯ
	// ─────────────────────────────────────────────────────
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bot started successfully!", bot.Self.UserName)

	// Настройка получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Канал задач для воркер-пула
	ch := make(chan models.Job, 100)

	// ─────────────────────────────────────────────────────
	// 3. ЗАПУСК КОМПОНЕНТОВ (ОРКЕСТРАЦИЯ)
	// ─────────────────────────────────────────────────────

	// Запускаем 4 воркера для параллельной обработки
	for i := 1; i <= 4; i++ {
		go worker.Worker(ch)
	}

	// Запускаем обработчик входящих сообщений
	// (он будет отправлять задачи в канал ch)
	handler.Start(updates, ch, bot)
}
