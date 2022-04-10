package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) CheckPostgresVersion(c *fiber.Ctx) error {
	v, err := h.ServiceLayer.CheckPostgresVersion()
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"Version": v,
	})
}
