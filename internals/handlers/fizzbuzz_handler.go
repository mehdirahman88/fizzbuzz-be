package fizzbuzzHandler

import (
	"github.com/gofiber/fiber/v2"
	fizzbuzzService "github.com/mehdirahman88/fizzbuzz-be/internals/services"
	"strconv"
)

func GetFizzBuzz(c *fiber.Ctx) error {
	counter, err := strconv.Atoi(c.Params("counter"))

	if err != nil {
		return c.JSON(fiber.Map{
			"isSuccess": false,
			"error":     "Counter must be Integer",
		})
	}

	return c.JSON(fiber.Map{
		"isSuccess": true,
		"data":      fizzbuzzService.CalcFizzBuzz(counter),
	})
}
