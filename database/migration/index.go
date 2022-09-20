package migration

import (
	"fmt"
	"go-gin-api/database"
	"go-gin-api/models/entities"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entities.Users{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
