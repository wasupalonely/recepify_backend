package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	JwtSecret   string
}

var AppConfig Config

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	AppConfig = Config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}