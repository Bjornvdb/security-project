package handlers

import (
	"security-project/database"
	"security-project/models"

	"github.com/gofiber/fiber/v2"
)

// GetUsers returns all the users
func GetUsers(c *fiber.Ctx) error {
	db := database.DB

	var results []models.UserDTO
	db.Table("users").Find(&results)

	return c.JSON(fiber.Map{
		"users": results,
	})
}

// GetUser returns the user
func GetUser(c *fiber.Ctx) error {
	db := database.DB

	var user models.User

	db.First(&user)

	t := user.CreatedAt

	return c.JSON(fiber.Map{
		"user": user,
		"t":    t,
	})
}
