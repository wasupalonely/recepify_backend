package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `json:"email" binding:"required,email"`
	ProfilePicture string `json:"profile_picture" binding:"nullable"`
	Bio            string `json:"bio" binding:"nullable"`
	Password       string `json:"password" binding:"required,min=6,max=64"`
	Username       string `json:"username" binding:"required"`
}
