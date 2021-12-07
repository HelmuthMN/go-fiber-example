package main

import (
	"github.com/HelmuthMN/go-fiber-example/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connection to the Database
	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Testing air")
		return err
	})
	app.Listen(":3000")
}
