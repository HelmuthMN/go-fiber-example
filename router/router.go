package router

import (
	noteRoutes "github.com/HelmuthMN/go-fiber-example/internal/routes/note"
	userRoutes "github.com/HelmuthMN/go-fiber-example/internal/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	//Setup the Note Routes
	noteRoutes.SetupNoteRoutes(api)
	userRoutes.SetupUserRoutes(api)
}
