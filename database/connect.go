package database

import (
	"fmt"
	"security-project/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Print("failed")
		panic("failed to connect database")
	}

	if !fiber.IsChild() {
		DB.AutoMigrate(&models.User{}, models.Todo{})
		fmt.Println("Database Migrated")
	}
}
