package routes

import "github.com/gin-gonic/gin"

func RouteInitialization(router *gin.RouterGroup) {
	userRoutes(router)
}
