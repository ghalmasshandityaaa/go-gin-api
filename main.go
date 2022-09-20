package main

import (
	"go-gin-api/database"
	"go-gin-api/database/migration"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	database.DBInitialization()
	migration.RunMigration()

	app.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"OK":      false,
			"message": "Ooppss! 404 api not found!",
		})
	})

	app.Run(":5500")
}
