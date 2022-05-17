package handlers

import (
	"fmt"
	"os"
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

func (h *CustomerHandler) SaveImage(c *fiber.Ctx) error {
	id := new(models.IdReg)
	if err := c.BodyParser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	validate = validator.New()
	if err := validate.Struct(id); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(500).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	path, err := h.handler.SaveImage(id, file.Filename)
	if err != nil {
		return c.Status(500).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := c.SaveFile(file, "."+path); err != nil {
		return c.Status(500).JSON(models.ErrorResp{
			Status:  false,
			Message: "Save File handler" + err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"Message": "Image uploaded successfully",
	})
}

func (h *CustomerHandler) DeleteImage(c *fiber.Ctx) error {
	image := new(models.Image)
	if err := c.BodyParser(image); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	validate = validator.New()
	if err := validate.Struct(image); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	// delete image from ./images
	err := os.Remove(fmt.Sprintf("./images/customer/%d/%s", image.Id, image.Name))
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	if err := h.handler.DeleteImage(&models.IdReg{image.Id}); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  201,
		"message": "Image deleted successfully",
	})
}

func (h *CustomerHandler) GmailCode(c *fiber.Ctx) error {
	email := new(models.EmailRequest)
	if err := c.BodyParser(email); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	validate = validator.New()
	if err := validate.Struct(email); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	code, err := h.handler.GmailCode(email)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
		"code":    code,
	})
}

func (h *CustomerHandler) Setter(c *fiber.Ctx) error {
	deal := new(models.Deal)
	if err := c.BodyParser(deal); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	validate = validator.New()
	if err := validate.Struct(deal); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.handler.Setter(deal); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"Status":  false,
			"Message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
	})
}

func (h *CustomerHandler) Getter(c *fiber.Ctx) error {
	id := new(models.ProductId)
	if err := c.BodyParser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	validate = validator.New()
	if err := validate.Struct(id); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	v, err := h.handler.Getter(id)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"Status":  false,
			"Message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
		"Value":   v,
	})
}