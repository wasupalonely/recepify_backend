package recipe

import (
	"github.com/wasupalonely/recepify/internal/models"
	"github.com/wasupalonely/recepify/pkg/db"
)

func CreateRecipe(recipe *models.Recipe) error {
	return db.DB.Create(recipe).Error
}

func GetAllRecipes() ([]models.Recipe, error) {
	var recipes []models.Recipe
	if err := db.DB.Preload("Categories").Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func GetRecipeByID(id string) (*models.Recipe, error) {
	var recipe models.Recipe
	if err := db.DB.Preload("Categories").First(&recipe, id).Error; err != nil {
		return nil, err
	}
	return &recipe, nil
}

func UpdateRecipe(recipe *models.Recipe) error {
	return db.DB.UpdateColumns(recipe).Error
}

func DeleteRecipe(id string) error {
	return db.DB.Delete(&models.Recipe{}, id).Error
}
