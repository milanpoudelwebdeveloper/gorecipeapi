package routes

import (
	"recipeapi/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	categoryController := controllers.CategoriesController{}
	categoryGroup := r.Group("/categories")
	{
		categoryGroup.GET("/", categoryController.GetCategories)
		categoryGroup.POST("/", categoryController.AddCategory)
	}
}
