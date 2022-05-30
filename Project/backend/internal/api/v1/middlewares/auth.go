package middlewares

import (
	"os"
	"secondChance/internal/models"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("secret")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(models.Resp{
				Status:  false,
				Message: err.Error(),
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
}
