package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wasupalonely/recepify/internal/auth"
	"github.com/wasupalonely/recepify/internal/category"
	"github.com/wasupalonely/recepify/internal/recipe"
	"github.com/wasupalonely/recepify/internal/user"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/login", auth.LoginHandler)
		api.POST("/register", auth.RegisterHandler)

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
				categoriesGroup.GET("/", category.GetCategoriesHandler)
				categoriesGroup.GET("name/:name", category.GetCategoryByNameHandler)
				categoriesGroup.GET("/:id", category.GetCategoryHandler)
				categoriesGroup.POST("/", category.CreateCategoryHandler)
				categoriesGroup.PUT("/:id", category.UpdateCategoryHandler)
				categoriesGroup.DELETE("/:id", category.DeleteCategoryHandler)
			}

			recipesGroup := authGroup.Group("/recipes")
			{
				recipesGroup.GET("/:id", recipe.GetRecipeHandler)
				recipesGroup.GET("/", recipe.GetRecipesHandler)
				recipesGroup.POST("/", recipe.CreateRecipeHandler)
				recipesGroup.PUT("/:id", recipe.UpdateRecipeHandler)
				recipesGroup.DELETE("/:id", recipe.DeleteRecipeHandler)
			}
		}
	}

	return r
}
