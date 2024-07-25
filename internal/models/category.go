package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Image string `json:"image"`
	Recipes []Recipe `gorm:"many2many:recipe_categories;" json:"recipes"`
}
