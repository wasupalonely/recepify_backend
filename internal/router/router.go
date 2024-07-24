package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wasupalonely/recepify/internal/auth"
	"github.com/wasupalonely/recepify/internal/user"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/login", auth.LoginHandler)
		api.POST("/register", auth.RegisterHandler)
	}

	authGroup := r.Group("/")
	authGroup.Use(auth.AuthMiddleware())
	{
		usersGroup := api.Group("/users")
		{
			usersGroup.GET("/:id", user.GetUserHandler)
			usersGroup.GET("/", user.GetUsersHandler)
			usersGroup.POST("/", user.CreateUserHandler)
			// TODO: Add UpdateUserHandler
			// usersGroup.PUT("/:id", user.UpdateUserHandler)
			usersGroup.DELETE("/:id", user.DeleteUserHandler)
		}
	}

	return r
}