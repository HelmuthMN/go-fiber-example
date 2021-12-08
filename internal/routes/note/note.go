package noteRoutes

import "github.com/gofiber/fiber/v2"

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	note.Post("/", func(c *fiber.Ctx) error {})
	note.Get("/", func(c *fiber.Ctx) error {})
	note.Get("/:noteId", func(c *fiber.Ctx) error {})
	note.Put("/:noteId", func(c *fiber.Ctx) error {})
	note.Delete("/:noteId", func(c *fiber.Ctx) error {})
}
