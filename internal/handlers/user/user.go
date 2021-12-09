package noteHandler

import (
	"github.com/HelmuthMN/go-fiber-example/database"
	"github.com/HelmuthMN/go-fiber-example/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	ID       uuid.UUID `json:"id"`
// 	Name     string    `json:"name"`
// 	Email    string    `json:"email" gorm:"unique"`
// 	Password string    `json:"password"`
// }

// func createUserResponse(userModel models.User) User {
// 	return User{ID: userModel.ID, Name: userModel.Name, Email: userModel.Email, Password: userModel.Password}
// }

func RegisterUser(c *fiber.Ctx) error {
	db := database.DB
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		ID:       uuid.New(),
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	// database.DB.Create(&user)
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// responseUser := createUserResponse(user)

	return c.Status(200).SendString("Registration completed")
}

// TODO: User login and  JWT
