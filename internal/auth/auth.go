package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wasupalonely/recepify/config"
	"github.com/wasupalonely/recepify/internal/user"
	"github.com/wasupalonely/recepify/internal/validations"
)

func generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JwtSecret))
}

func LoginHandler(c *gin.Context) {
	var loginData validations.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := user.GetByIdentifier(loginData.Identifier)
	if err != nil || !user.CheckPasswordHash(loginData.Password, u.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid identifier or password"})
		return
	}

	token, err := generateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate the token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": gin.H{"username": u.Username, "profilePicture": u.ProfilePicture}})
}

func RegisterHandler(c *gin.Context) {
	var registerData validations.RegistrationData
	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := user.HashPassword(registerData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	newUser := user.User{
		Username: registerData.Username,
		Password: hashedPassword,
		Email:    registerData.Email,
	}

	if err := user.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
