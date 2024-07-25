package recipe

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wasupalonely/recepify/internal/models"
	"github.com/wasupalonely/recepify/pkg/db"
)

func CreateRecipeHandler(c *gin.Context) {
	var recipeInput struct {
		Title       string              `json:"title" binding:"required"`
		Description string              `json:"description" binding:"required"`
		Ingredients []models.Ingredient `json:"ingredients" binding:"required"`
		Steps       []models.Step       `json:"steps"`
		UserID      uint                `json:"user_id" binding:"required"`
		CategoryIDs []uint              `json:"category_ids"`
	}

	if err := c.ShouldBindJSON(&recipeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ingredientsJSON, err := json.Marshal(recipeInput.Ingredients)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse ingredients"})
		return
	}

	recipe := models.Recipe{
		Title:       recipeInput.Title,
		Description: recipeInput.Description,
		Ingredients: ingredientsJSON,
		UserID:      recipeInput.UserID,
	}

	tx := db.DB.Begin()

	if err := tx.Create(&recipe).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create recipe"})
		return
	}

	for i := range recipeInput.Steps {
		recipeInput.Steps[i].RecipeID = recipe.ID
	}
	if err := tx.Create(&recipeInput.Steps).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create steps"})
		return
	}

	if err := tx.Model(&recipe).Association("Steps").Append(&recipeInput.Steps); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not associate steps"})
		return
	}

	if len(recipeInput.CategoryIDs) > 0 {
		var categories []models.Category
		if err := tx.Find(&categories, recipeInput.CategoryIDs).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"error": "Could not find categories"})
			return
		}
		if err := tx.Model(&recipe).Association("Categories").Append(&categories); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not associate categories"})
			return
		}
	}

	tx.Commit()

	c.JSON(http.StatusOK, recipe)
}


func GetRecipesHandler(c *gin.Context) {
	recipes, err := GetAllRecipes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get recipes"})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

func GetRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	recipe, err := GetRecipeByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get recipe"})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

func UpdateRecipeHandler(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := UpdateRecipe(&recipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update recipe"})
		return
	}
}

func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteRecipe(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete recipe"})
		return
	}
}
