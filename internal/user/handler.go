package user

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wasupalonely/recepify/internal/models"
	"github.com/wasupalonely/recepify/internal/uploads/cloudinary"
	"github.com/wasupalonely/recepify/internal/validations"
)

func UpdateProfilePictureHandler(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := "./temp/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	defer os.Remove(filePath)

	imageURL, err := cloudinary.UploadImage(filePath, "profile-pictures")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	userIDConv, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := UpdateUserProfilePicture(uint(userIDConv), imageURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile picture"})
		return
	}

	// Devolver la URL de la imagen subida
	c.JSON(http.StatusOK, gin.H{"url": imageURL})
}


func CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = hashedPassword

	if err := CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsersHandler(c *gin.Context) {
	users, err := GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserHandler(c *gin.Context) {
	userID := c.Param("id")
	fmt.Println("userID: ", userID)
	user, err := GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetUserByUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	user, err := GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUserHandler(c *gin.Context) {
	userID := c.Param("id")
	var user validations.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIDConv, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(userIDConv)
	if err := UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")
	fmt.Println("userID: ", userID)
	if err := DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
