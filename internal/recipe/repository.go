package recipe

import (
	"github.com/wasupalonely/recepify/internal/models"
	"github.com/wasupalonely/recepify/pkg/db"
)

func CreateRecipe(recipe *models.Recipe) error {
	return db.DB.Create(recipe).Error
}

func GetAllRecipes(limit int, offset int) ([]models.Recipe, error) {
	var recipes []models.Recipe
	if err := db.DB.Debug().Preload("Steps").Preload("Categories").Limit(limit).Offset(offset).Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func GetRecipeByID(id string) (*models.Recipe, error) {
	var recipe models.Recipe
	if err := db.DB.Debug().Preload("Steps").Preload("Categories").First(&recipe, id).Error; err != nil {
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

func GetRecipesByUserID(userID string, limit int, offset int) ([]models.Recipe, error) {
	var recipes []models.Recipe
	if err := db.DB.Preload("Categories").Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

// Steps section

func CreateSteps(step []*models.Step) error {
	return db.DB.Create(step).Error
}
