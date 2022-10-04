package services

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"test-api/repositories"
)

type Success struct {
	Success bool     `json:"success"`
	Data    struct{} `json:"data"`
}

func GetUser(app *fiber.Ctx) error {
	user := new(Success)
	user.Success = true

	return app.JSON(user)
}

func CreateUser(app *fiber.Ctx) error {
	var user repositories.User

	err := app.BodyParser(&user)

	if err != nil {
		log.Fatal(err)
	}

	// Create user.
	repositories.CreateUser(&user)

	return app.JSON(user)
}
