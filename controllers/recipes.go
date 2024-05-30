package controllers

import (
	"fmt"
	"net/http"
	"recipeapi/db"
	"recipeapi/models"

	"github.com/gin-gonic/gin"
)

type RecipeController struct{}

// @Summary Get all recipes
// @Description Get a list of all recipes
// @Tags recipes
// @Produce json
// @Success 200 {array} models.Recipe
// @Failure 500 {object} gin.H{"message": string}
// @Router /recipes [get]

func (ctrl RecipeController) GetRecipes(c *gin.Context) {
	query := `SELECT * FROM recipe`
	rows, error := db.DB.Query(query)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while fetching recipes",
		})
		return
	}
	var recipes []models.Recipe
	for rows.Next() {
		var recipe models.Recipe
		err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.Country, &recipe.Instructions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while fetching recipes",
			})
			return
		}
		recipes = append(recipes, recipe)

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "All recipes",
		"data":    recipes,
	})
}

func (ctrl RecipeController) AddRecipes(c *gin.Context) {
	var newRecipe models.Recipe
	query := `INSERT INTO recipe(title, description, country, instructions) VALUES ($1, $2, $3, $4)`
	err := c.BindJSON(&newRecipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong while parsing body",
		})
		return
	}
	_, err = db.DB.Exec(query, newRecipe.Title, newRecipe.Description, newRecipe.Country, newRecipe.Instructions)
	if err != nil {
		fmt.Println("error here is", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while inserting recipe",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Recipe added successfully",
	})
}
