package controllers

import (
	"net/http"
	"recipeapi/models"

	"github.com/gin-gonic/gin"
)

type RecipeController struct{}

func (ctrl RecipeController) GetRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "All recipes",
	})
}

func (ctrl RecipeController) AddRecipes(c *gin.Context) {
	var newRecipe models.Recipe
	err := c.BindJSON(&newRecipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong while parsing body",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "New recipe added successfully",
		"data":    newRecipe,
	})

}
