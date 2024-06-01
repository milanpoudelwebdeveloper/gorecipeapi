package routes

import (
	"recipeapi/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userController := controllers.UserController{}
	userGroup := router.Group("/users")
	{
		userGroup.PUT("/update", userController.UpdateProfile)
		userGroup.PUT("/change/password", userController.ChangePassword)
	}
}
