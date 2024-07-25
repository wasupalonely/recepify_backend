package validations

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `json:"name" binding:"omitempty"`
	Image string `json:"image" binding:"omitempty"`
}