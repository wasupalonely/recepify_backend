package user

import "github.com/wasupalonely/recepify/pkg/db"

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetByIdentifier(identifier string) (*User, error) {
	var user User
	if err := db.DB.Where("username = ? OR email = ?", identifier, identifier).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(userID uint) (*User, error) {
	var user User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) error {
	return db.DB.Create(user).Error
}

func UpdateUser(user *User) error {
	return db.DB.Save(user).Error
}

func DeleteUser(userID uint) error {
	return db.DB.Delete(&User{}, userID).Error
}

