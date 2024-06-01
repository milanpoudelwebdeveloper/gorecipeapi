package adminRoutes

import (
	"recipeapi/controllers/admin"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	adminController := admin.AdminController{}
	admin := r.Group("/admin")
	{
		admin.GET("/users", adminController.GetUsers)
		admin.POST("/users", adminController.CreateUser)
	}
}
