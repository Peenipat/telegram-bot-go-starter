package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBUrl         string
	TelegramToken string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	AppConfig = Config{
		Port:          getEnv("PORT", "8080"),
		DBUrl:         getEnv("DATABASE_URL", os.Getenv("DB_URL")),
		TelegramToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
