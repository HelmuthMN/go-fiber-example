package main

import (
	"github.com/HelmuthMN/go-fiber-example/database"
	"github.com/HelmuthMN/go-fiber-example/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connection to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)
	app.Listen(":3000")
}
