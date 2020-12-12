package routes

import (
	"errors"
	"fmt"
	"security-project/database"
	"security-project/handlers"
	"security-project/models"
	"security-project/session"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"golang.org/x/crypto/bcrypt"
)

// Middleware to ensure the requestor is logged in
func authMW(c *fiber.Ctx) error {
	sess, err := session.Sessions.Get(c)
	if err != nil {
		return err
	}

	uid := sess.Get("userID")
	if uid == nil {
		return errors.New("u bent niet ingelogd")
	}
	c.Locals("userID", uid)

	return c.Next()
}

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	v1 := api.Group("/v1")

	// HTML routes
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{
			"csrf": c.Locals("csrf"),
		})
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", fiber.Map{})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		db := database.DB
		sess, err := session.Sessions.Get(c)
		if err != nil {
			panic(err)
		}

		id := sess.Get("userID")

		if id != nil {
			var userGORM models.User

			db.Preload("Todos").First(&userGORM, id)

			return c.Render("index", fiber.Map{
				"User": userGORM,
			})
		}

		return c.Render("index", fiber.Map{})
	})

	app.Get("/todos/add", authMW, func(c *fiber.Ctx) error {
		sess, err := session.Sessions.Get(c)
		if err != nil {
			panic(err)
		}
		sess.Set("lang", c.Get("Accept-Language"))
		sess.Save()
		return c.Render("add", fiber.Map{
			"csrf": c.Locals("csrf"),
		})
	})

	app.Post("/todos/add", authMW, func(c *fiber.Ctx) error {
		db := database.DB
		todo := new(models.TodoDTO)

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		fmt.Print(c.Locals("userID"))

		var userGORM models.User

		db.Where("ID = ?", c.Locals("userID")).First(&userGORM)

		db.Model(&userGORM).Association("Todos").Append(&models.Todo{Name: todo.Name, Body: todo.Body})

		return c.Redirect("/")
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		db := database.DB
		user := new(models.UserDTO)

		if err := c.BodyParser(user); err != nil {
			return err
		}

		fmt.Println(user.Username)

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// Create user

		userGORM := models.User{Username: user.Username, Password: string(hash)}
		db.Create(&userGORM)
		if err != nil {
			return err
		}

		return c.Redirect("/login")
	})


	app.Post("/login", func(c *fiber.Ctx) error {
		db := database.DB
		body := new(models.UserDTO)
		// Decode body
		if err := c.BodyParser(body); err != nil {
			return err
		}

		// Find user
		var userGORM models.User

		db.Where("username = ?", body.Username).First(&userGORM)


		// Validate password
		err := bcrypt.CompareHashAndPassword([]byte(userGORM.Password), []byte(body.Password))
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.SendString("Invalid password")
		}

		sess, err := session.Sessions.Get(c)
		if err != nil {
			panic(err)
		}
		sess.Set("userID", userGORM.ID)
		sess.Save()

		return c.Redirect("/")
	})

















	user := v1.Group("/users")
	user.Get("/", handlers.GetUsers)
	user.Get("/:id", handlers.GetUser)
	user.Put("/", handlers.GetUsers)

	session := v1.Group("/session")
	session.Get("/get", handlers.GetSession)
	session.Get("/set", handlers.SetSession)

	auth := v1.Group("/auth")

	auth.Post("/login", handlers.Login)
}
