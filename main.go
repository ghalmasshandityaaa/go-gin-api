package main

import (
	"go-gin-api/database"
	"go-gin-api/database/migration"
	"go-gin-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	database.DBInitialization()
	migration.RunMigration()

	router := app.Group("/api/v1")
	routes.RouteInitialization(router)

	app.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"OK":      false,
			"message": "Ooppss! 404 api not found!",
		})
	})

	app.Run(":5500")
}
