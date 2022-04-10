package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"secondChance/internal/models"
	"time"
)

func (h *Handler) GetAllProduct(c *fiber.Ctx) (err error) {
	products, err := h.ServiceLayer.GetAllProduct()
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

func (h *Handler) GetProduct(c *fiber.Ctx) (err error) {
	id := new(models.IdReg)
	if err := c.QueryParser(id); err != nil {
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

	product, err := h.ServiceLayer.GetProduct(id)
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
		"product": product,
	})
}

func (h *Handler) SoldProduct(c *fiber.Ctx) (err error) {
	var id models.IdReg
	if err := c.QueryParser(&id); err != nil {
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

	if err := h.ServiceLayer.SoldProduct(time.Now(), id); err != nil {
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

func (h *Handler) CreateProduct(c *fiber.Ctx) (err error) {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(product); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.ServiceLayer.CreateProduct(product); err != nil {
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

func (h *Handler) UpdateProduct(c *fiber.Ctx) (err error) {
	id := new(models.IdReg)
	if err := c.QueryParser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	validate = validator.New()
	if err := validate.Struct(id); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}

	productDB, err := h.ServiceLayer.GetProduct(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResp{
				Status:  false,
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusConflict).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	productReq := new(models.Product)
	if err := c.BodyParser(productReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	//Todo Status Code
	product := newProduct(productReq, productDB)
	if err := h.ServiceLayer.UpdateProduct(id, product); err != nil {
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

// TODO Valid
func newProduct(productReq, productDB *models.Product) *models.Product {
	if productReq.Price != 0 {
		productDB.Price = productReq.Price
	}
	if productReq.Name != "" {
		productDB.Name = productReq.Name
	}
	if productReq.Description != "" {
		productDB.Description = productReq.Description
	}
	return productDB
}
