package handlers

import (
	"security-project/session"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// SetSession returns all the users
func SetSession(c *fiber.Ctx) error {
	sess, err := session.Sessions.Get(c)
	if err != nil {
		panic(err)
	}
	defer sess.Save()
	sess.Set("test", utils.UUID())

	return c.JSON(fiber.Map{
		"users": "succes",
	})
}

// GetSession returns the user
func GetSession(c *fiber.Ctx) error {
	store, err := session.Sessions.Get(c)
	if err != nil {
		panic(err)
	}
	defer store.Save()
	return c.JSON(fiber.Map{
		"session": store.Get("test"),
	})
}
