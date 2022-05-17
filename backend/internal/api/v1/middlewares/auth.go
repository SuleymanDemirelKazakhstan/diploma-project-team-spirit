package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"os"
	"secondChance/internal/models"
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
			JSON(models.ErrorResp{
				Status:  false,
				Message: err.Error(),
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
}
