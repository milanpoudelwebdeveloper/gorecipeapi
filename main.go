package main

import (
	"recipeapi/db"
	_ "recipeapi/docs"
	"recipeapi/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Recipe API
// @version 1.0
// @description This is a sample server for a recipe API.
// @contact.name Milan Poudel
// @contact.url https://github.com/milanpoudelwebdeveloper
// @host localhost:8080
// @BasePath /

func main() {
	db.Init()
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.UserRoutes(r)
	routes.RecipeRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
