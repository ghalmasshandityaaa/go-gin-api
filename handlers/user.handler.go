package handlers

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/database"
	"go-gin-api/models/entities"
)

func GetUsers(ctx *gin.Context) {
	var users []entities.Users

	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		ctx.JSON(200, gin.H{
			"OK":      false,
			"message": result.Error,
		})
	}

	ctx.JSON(200, gin.H{
		"OK":      true,
		"message": "Successfully get all users",
		"data":    users,
	})
}

func GetUserById(ctx *gin.Context) {
	var user entities.Users
	userId := ctx.Param("id")
	if userId == "" {
		ctx.JSON(400, gin.H{
			"OK":      false,
			"message": "userId is required",
		})
		return
	}

	result := database.DB.Debug().First(&user, "id = ?", userId)
	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"OK":      false,
			"message": "User not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"OK":      true,
		"message": "Successfully get user",
		"data":    user,
	})
}
