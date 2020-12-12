package handlers

import (
	"security-project/session"
	"security-project/utils"

	"github.com/gofiber/fiber/v2"
)

type userLoginDTO struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// Login Will try to attempt to login a user
func Login(c *fiber.Ctx) error {
	user := new(userLoginDTO)

	if err := c.BodyParser(user); err != nil {
		return utils.BadRequest(c)
	}

	if user.Email == "test@test.be" && user.Password == "test1234" {
		store, err := session.Sessions.Get(c)
		if err != nil {
			panic(err)
		}
		defer store.Save()
		store.Set("userid", user.Email)

		//hash, err := argon2id.CreateHash("pa$$word", argon2id.DefaultParams)

		/* 		if err != nil {
			log.Fatal(err)
		} */

		return c.JSON(fiber.Map{
			"user": user,
			//"hash": hash,
		})
	}

	return utils.BadRequest(c)
}
