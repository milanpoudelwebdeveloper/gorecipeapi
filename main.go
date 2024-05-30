package main

import (
	"recipeapi/db"
	"recipeapi/routes"

	"github.com/gin-gonic/gin"
)

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
	r.Run()
}
