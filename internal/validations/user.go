package validations

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email          string `json:"email" binding:"omitempty,email"`
    Bio            string `json:"bio" binding:"omitempty"`
    Username       string `json:"username" binding:"omitempty"`
}
