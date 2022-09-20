package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/handlers"
)

func userRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("/", handlers.GetUsers)
		users.GET("/:id", handlers.GetUserById)
	}
}
