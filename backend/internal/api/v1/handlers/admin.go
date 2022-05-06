package handlers

import (
	"fmt"
	"os"
	"secondChance/internal/models"
	"secondChance/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

type AdminHandler struct {
	handler services.Admin
}

func NewAdminHandler(service services.Admin) *AdminHandler {
	return &AdminHandler{handler: service}
}

func (h *AdminHandler) Create(c *fiber.Ctx) (err error) {
	a := new(models.Owner)
	if err := c.BodyParser(a); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(a); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.handler.Create(a); err != nil {
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

func (h *AdminHandler) Delete(c *fiber.Ctx) (err error) {
	a := new(models.IdReg)
	if err := c.BodyParser(a); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(a); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.handler.Delete(a); err != nil {
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

func (h *AdminHandler) Get(c *fiber.Ctx) (err error) {
	a := new(models.IdReg)
	if err := c.BodyParser(a); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(a); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	user, err := h.handler.Get(a)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResp{
				Status:  false,
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
		"user":    user,
	})
}

func (h *AdminHandler) GetAll(c *fiber.Ctx) (err error) {
	users, err := h.handler.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
		"users":   users,
	})
}

func (h *AdminHandler) Update(c *fiber.Ctx) (err error) {
	user := new(models.Owner)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.handler.Update(user); err != nil {
		return c.JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
	})
}

func (h *AdminHandler) Login(c *fiber.Ctx) (err error) {
	param := new(models.LoginInput)
	if err := c.BodyParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	t, err := h.handler.Login(param)
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

func (h *AdminHandler) SaveImage(c *fiber.Ctx) error {
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

	if err := c.SaveFile(file, path); err != nil {
		return c.Status(500).JSON(models.ErrorResp{
			Status:  false,
			Message: "Save File handler" + err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"Message": "Image uploaded successfully",
	})
}

func (h *AdminHandler) DeleteImage(c *fiber.Ctx) error {
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
	err := os.Remove(fmt.Sprintf("./images/shop/%d/%s", image.Id, image.Name))
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
