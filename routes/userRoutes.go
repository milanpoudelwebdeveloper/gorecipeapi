package routes

import (
	"recipeapi/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userController := controllers.UserController{}
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", userController.GetUsers)
		userGroup.POST("/", userController.AddUser)
	}
}
