package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fizzbuzzRoute "github.com/mehdirahman88/fizzbuzz-be/internals/routes"
)

func SetupRoutes(app *fiber.App) {
	// Group endpoints and prefix /api
	api := app.Group("/api", logger.New())

	// Setup fizzbuzz route
	fizzbuzzRoute.SetupNoteRoutes(api)
}
