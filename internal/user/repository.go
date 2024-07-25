package user

import (
	"github.com/wasupalonely/recepify/internal/models"
	"github.com/wasupalonely/recepify/internal/validations"
	"github.com/wasupalonely/recepify/pkg/db"
)

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetByIdentifier(identifier string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("username = ? OR email = ?", identifier, identifier).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func UpdateUser(user *validations.User) error {
	return db.DB.UpdateColumns(user).Error
}

func DeleteUser(userID string) error {
	return db.DB.Delete(&models.User{}, userID).Error
}

