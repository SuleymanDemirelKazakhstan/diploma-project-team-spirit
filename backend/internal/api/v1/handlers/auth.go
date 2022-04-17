package handlers

import (
	"github.com/gofiber/fiber/v2"
	"secondChance/internal/models"
)

// Login get user and password
func (h *Handler) LoginCustomer(c *fiber.Ctx) (err error) {
	var param models.LoginInput
	if err := c.BodyParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	t, id, err := h.ServiceLayer.LoginCustomer(&param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":      "success",
		"token":       t,
		"customer_id": id,
	})
}

func (h *Handler) LoginOwner(c *fiber.Ctx) (err error) {
	param := new(models.LoginInput)
	if err := c.BodyParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	t, err := h.ServiceLayer.LoginOwner(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"token":  t,
	})
}
