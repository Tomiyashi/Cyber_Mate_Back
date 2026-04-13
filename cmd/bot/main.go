package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"CyberMate_Back/internal/config"
	"CyberMate_Back/internal/handler"
	"CyberMate_Back/internal/models"
	"CyberMate_Back/internal/worker"
	"CyberMate_Back/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env not found, using env variables")
	}

	// Инициализация конфига
	cfg := config.New()

	// Инициализация БД (pgxpool)
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ Failed to create DB pool: %v", err)
	}
	defer pool.Close()

	// Проверка подключения к БД
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("❌ Failed to ping DB: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL")

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("❌ Failed to create bot: %v", err)
	}
	log.Printf("✅ Authorized as @%s", bot.Self.UserName)

	// Инициализация репозитория
	userRepo := repositories.NewUserRepository(pool)

	// Канал для задач
	jobs := make(chan models.Job, 100)
	workerCount := runtime.NumCPU() * 2
	if workerCount < 4 {
		workerCount = 4
	}
	for i := 0; i < workerCount; i++ {
		go worker.Worker(jobs)
	}

	// Хендлер
	h := handler.New(bot, jobs, userRepo, cfg.MiniAppURL)

	// Запуск бота
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	log.Println("🚀 Bot started...")

	for update := range updates {
		h.Handle(update)
	}
}
