package main

import (
	"github.com/HelmuthMN/go-fiber-example/database"
	"github.com/HelmuthMN/go-fiber-example/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connection to the Database
	database.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Setup the router
	router.SetupRoutes(app)
	app.Listen(":3000")
}
