package routes

import (
	"recipeapi/controllers"

	"github.com/gin-gonic/gin"
)

func RecipeRoutes(r *gin.Engine) {
	recipeController := controllers.RecipeController{}
	recipeGroup := r.Group("/recipes")
	{
		recipeGroup.GET("/", recipeController.GetRecipes)
		recipeGroup.POST("/", recipeController.AddRecipes)
		recipeGroup.GET("/:id", recipeController.GetRecipeDetails)
		recipeGroup.PUT("/:id", recipeController.UpdateRecipe)
		recipeGroup.DELETE("/:id", recipeController.DeleteRecipe)
	}
}
