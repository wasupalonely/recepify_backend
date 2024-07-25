package models

import "gorm.io/gorm"

type Step struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image"`
	RecipeID    uint   `json:"recipe_id" binding:"required"`
}