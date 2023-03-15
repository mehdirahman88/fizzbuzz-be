package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mehdirahman88/fizzbuzz-be/router"
	"log"
)

func main() {
	// Start a new fiber app
	app := fiber.New()
	app.Use(recover.New())

	// Set up the router
	router.SetupRoutes(app)

	// Listen on PORT 5001
	log.Fatal(app.Listen(":5001"))
}
