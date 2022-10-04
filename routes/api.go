package routes

import (
	"github.com/gofiber/fiber/v2"
	"test-api/services"
)

func RegisterApiRoutes(app *fiber.App) {
	version1(app)
}

func version1(app *fiber.App) {
	api := app.Group("api/")
	v1 := api.Group("v1/")

	// Users route.
	v1.Post("/users", services.CreateUser)
	v1.Get("/users/:id<int>", services.GetUser)
}
