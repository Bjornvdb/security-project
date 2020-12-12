package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// BadRequest returns if it is a bad request
func BadRequest(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{
		"status":    400,
		"error":     "Bad request",
		"timestamp": time.Now(),
	})
}
