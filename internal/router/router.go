package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wasupalonely/recepify/internal/auth"
	"github.com/wasupalonely/recepify/internal/category"
	"github.com/wasupalonely/recepify/internal/recipe"
	"github.com/wasupalonely/recepify/internal/user"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api/v1")
	{
		// Root endpoint
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Recepify"})
		})

		// Auth
		api.POST("/login", auth.LoginHandler)
		api.POST("/register", auth.RegisterHandler)

		// No auth recipes
		recipesGroup := api.Group("/recipes")
		{
			recipesGroup.GET("/:id", recipe.GetRecipeHandler)
			recipesGroup.GET("/", recipe.GetRecipesHandler)
		}

		// No auth categories
		categoriesGroup := api.Group("/categories")
		{
			categoriesGroup.GET("/", category.GetCategoriesHandler)
		}

		authGroup := api.Group("/")
		authGroup.Use(auth.AuthMiddleware())
		{
			usersGroup := authGroup.Group("/users")
			{
				usersGroup.GET("/:id", user.GetUserHandler)
				usersGroup.GET("/", user.GetUsersHandler)
				usersGroup.POST("/", user.CreateUserHandler)
				usersGroup.PUT("/:id", user.UpdateUserHandler)
				usersGroup.DELETE("/:id", user.DeleteUserHandler)
				usersGroup.POST("/profile-picture/:id", user.UpdateProfilePictureHandler)
			}

			categoriesGroup := authGroup.Group("/categories")
			{
				categoriesGroup.GET("name/:name", category.GetCategoryByNameHandler)
				categoriesGroup.GET("/:id", category.GetCategoryHandler)
				categoriesGroup.POST("/", category.CreateCategoryHandler)
				categoriesGroup.PUT("/:id", category.UpdateCategoryHandler)
				categoriesGroup.DELETE("/:id", category.DeleteCategoryHandler)
			}

			recipesGroup := authGroup.Group("/recipes")
			{
				recipesGroup.POST("/", recipe.CreateRecipeHandler)
				recipesGroup.PUT("/:id", recipe.UpdateRecipeHandler)
				recipesGroup.DELETE("/:id", recipe.DeleteRecipeHandler)
			}
		}
	}

	return r
}
