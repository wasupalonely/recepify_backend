package main

import (
	"log"

	"github.com/wasupalonely/recepify/config"
	"github.com/wasupalonely/recepify/internal/models"
	"github.com/wasupalonely/recepify/internal/router"
	"github.com/wasupalonely/recepify/internal/uploads/cloudinary"
	"github.com/wasupalonely/recepify/pkg/db"
)

func main() {
	config.InitConfig()

	cloudinary.Init(
        config.AppConfig.CloudinaryCloudName,
		config.AppConfig.CloudinaryApiKey,
		config.AppConfig.CloudinaryApiSecret,
    )

	if err := db.Init(); err != nil {
		panic(err)
	}

	// TODO: AUTOMIGRATE MODELS
	db.GetDB().AutoMigrate(&models.User{}, &models.Category{}, &models.Recipe{}, &models.Step{})

	r := router.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
