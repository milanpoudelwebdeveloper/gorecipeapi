package routes

import (
	"recipeapi/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	authController := controllers.AuthController{}
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/register", authController.Register)
		authGroup.PUT("/verify", authController.VerifyAccount)
		authGroup.POST("/forgot/password", authController.ForgotPassword)
		authGroup.PUT("/reset/password", authController.ResetPassword)
		authGroup.GET("/resend/verification", authController.ResendVerification)
		authGroup.GET("/logout", authController.Logout)
	}
}
