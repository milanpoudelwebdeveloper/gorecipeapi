package controllers

import (
	"fmt"
	"net/http"
	"recipeapi/db"
	"recipeapi/models"

	"github.com/gin-gonic/gin"
)

type RecipeController struct{}

// GetRecipes godoc
// @Summary get all recipes
// @Schemes
// @Description get all recipes
// @Tags recipes
// @Accept json
// @Produce json
// @Success 200 {object} models.Recipe
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
	defer rows.Close()
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

// AddRecipes godoc
// @Summary Post a recipe
// @Schemes
// @Description Post a recipe
// @Tags recipes
// @Accept json
// @Produce json
// @Param recipe body models.Recipe true "Recipe object that needs to be added"
// @Success 201 {string} string "Recipe added successfully"
// @Router /recipes [post]
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

func (ctrl RecipeController) GetRecipeDetails(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get recipe details",
	})
}

func (ctrl RecipeController) UpdateRecipe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update recipe",
	})
}

func (ctrl RecipeController) DeleteRecipe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete recipe",
	})
}
