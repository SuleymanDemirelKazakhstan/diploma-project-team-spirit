package handlers

import (
	"secondChance/internal/models"
	"secondChance/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	handler services.Customer
}

func NewCustomerHandler(s services.Customer) *CustomerHandler {
	return &CustomerHandler{handler: s}
}

// CreateUser new user
func (h *CustomerHandler) SignUp(c *fiber.Ctx) error {
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
	if err := h.handler.Create(user); err != nil {
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

func (h *CustomerHandler) GetOrder(c *fiber.Ctx) error {
	id := new(models.IdReg)
	if err := c.BodyParser(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(id); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	products, err := h.handler.GetOrder(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":   true,
		"message":  "success",
		"products": products,
	})
}

func (h *CustomerHandler) Login(c *fiber.Ctx) (err error) {
	var param models.LoginInput
	if err := c.BodyParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	t, id, err := h.handler.Login(&param)
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

func (h *CustomerHandler) Buy(c *fiber.Ctx) error {
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(order); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.handler.CreateOrder(order); err != nil {
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
