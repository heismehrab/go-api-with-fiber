package main

import (
	"log"

	"test-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Register routes.
	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(":30000"))
}
