package category

import (
	"github.com/wasupalonely/recepify/internal/models"
	"github.com/wasupalonely/recepify/internal/validations"
	"github.com/wasupalonely/recepify/pkg/db"
)

func CreateCategory(category *models.Category) error {
	return db.DB.Create(category).Error
}

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := db.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(id string) (*models.Category, error) {
	var category models.Category
	if err := db.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func GetCategoryByName(name string) (*models.Category, error) {
	var category models.Category
	if err := db.DB.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func UpdateCategory(category *validations.Category) error {
	return db.DB.UpdateColumns(category).Error
}

func DeleteCategory(id string) error {
	return db.DB.Delete(&models.Category{}, id).Error
}