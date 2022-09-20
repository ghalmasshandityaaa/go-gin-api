package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInitialization() {
	var err error

	dsn := "host=localhost user=root password=root dbname=gogin port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Database connected")
}
