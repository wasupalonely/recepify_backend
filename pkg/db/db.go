package db

import (
	"fmt"

	"github.com/wasupalonely/recepify/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	var err error

	dsn := config.AppConfig.DatabaseUrl
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Failed to connect database: %w", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return DB
}