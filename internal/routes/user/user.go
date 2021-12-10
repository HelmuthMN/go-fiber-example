package userRoutes

import (
	userHandler "github.com/HelmuthMN/go-fiber-example/internal/handlers/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")
	user.Post("/register", userHandler.RegisterUser)
	user.Post("/login", userHandler.Login)
	user.Get("/", userHandler.User)
	user.Post("/logout", userHandler.Logout)

}
