package main

import (
	"security-project/database"
	"security-project/routes"
	"security-project/session"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/template/html"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Prefork: false,
		Views: engine,
	})

	database.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowCredentials: true,
		MaxAge:           240,
	}))

	session.MakeSession()

	app.Use(helmet.New())

/* 	app.Use(csrf.New(csrf.Config{
		CookieDomain: "bjorn.be",
		ContextKey: "csrf",
		KeyLookup: "form:csrf",
	}))  */

	routes.SetupRoutes(app)

	app.Listen(":3000")

}
