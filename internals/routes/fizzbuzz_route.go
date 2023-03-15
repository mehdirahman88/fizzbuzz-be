package fizzbuzzRoute

import (
	"github.com/gofiber/fiber/v2"
	fizzbuzzHandler "github.com/mehdirahman88/fizzbuzz-be/internals/handlers"
)

func SetupNoteRoutes(router fiber.Router) {
	fizzbuzz := router.Group("/fizzbuzz")

	// Get FizzBuzz for counter
	fizzbuzz.Get("/:counter<int>", fizzbuzzHandler.GetFizzBuzz)
}
