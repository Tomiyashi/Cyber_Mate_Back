package config

import "os"

type Config struct {
	TelegramToken string
	DatabaseURL   string
	MiniAppURL    string
}

func New() *Config {
	return &Config{
		TelegramToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://localhost:5432/cybermate?sslmode=disable"),
		MiniAppURL:    getEnv("MINI_APP_URL", "https://t.me/CyberMate_bot"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
