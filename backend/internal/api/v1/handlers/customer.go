package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"secondChance/internal/models"
)

// CreateUser new user
func (h *Handler) SignUp(c *fiber.Ctx) error {
	user := new(models.Customer)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	hash, err := HashPassword(user.Password)
	if err!= nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	user.Password = hash
	if err := h.ServiceLayer.CreateCustomer(user);err!=nil{
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
	})
}
