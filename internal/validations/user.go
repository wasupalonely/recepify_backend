package validations

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string
	ProfilePicture string
	Bio            string
	Password       string
	Username       string
}

