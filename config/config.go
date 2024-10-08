package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	JwtSecret   string
	CloudinaryCloudName string
	CloudinaryApiKey string
	CloudinaryApiSecret string
}

var AppConfig Config

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	AppConfig = Config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
		CloudinaryCloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
		CloudinaryApiKey: os.Getenv("CLOUDINARY_API_KEY"),
		CloudinaryApiSecret: os.Getenv("CLOUDINARY_API_SECRET"),
	}
}