package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Ingredient struct {
	Name     string `json:"name" binding:"required"`
	Quantity string `json:"quantity" binding:"required"`
}

type Recipe struct {
	gorm.Model
	Title       string          `json:"title" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Image		string 			`json:"image"`
	Ingredients json.RawMessage `json:"ingredients" gorm:"type:jsonb" binding:"required"`
	UserID      uint            `json:"user_id" binding:"required"`
	Steps       []Step          `json:"steps"`
	Categories  []Category      `gorm:"many2many:recipe_categories;" json:"categories"`
	CategoryIDs []uint          `json:"category_ids" gorm:"-"`
}
