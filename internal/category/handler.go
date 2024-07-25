package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wasupalonely/recepify/internal/models"
	"github.com/wasupalonely/recepify/internal/validations"
)

func CreateCategoryHandler(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create category"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func GetCategoriesHandler(c *gin.Context) {
	categories, err := GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func GetCategoryHandler(c *gin.Context) {
	id := c.Param("id")
	category, err := GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get category"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func GetCategoryByNameHandler(c *gin.Context) {
	name := c.Param("name")
	category, err := GetCategoryByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get category"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func UpdateCategoryHandler(c *gin.Context) {
	var category validations.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := UpdateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update category"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func DeleteCategoryHandler(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
