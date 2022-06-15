package handlers

import (
	"fmt"
	"os"
	"secondChance/internal/models"
	"secondChance/internal/services"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type OwnerHandler struct {
	handler services.Shop
}

func NewOwnerHandler(s services.Shop) *OwnerHandler {
	return &OwnerHandler{handler: s}
}

func (h *OwnerHandler) GetAll(c *fiber.Ctx) (err error) {
	products, err := h.handler.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
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

func (h *OwnerHandler) Get(c *fiber.Ctx) (err error) {
	id := new(models.IdReg)
	if err := c.QueryParser(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(id); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	product, shop, err := h.handler.Get(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.Status(fiber.StatusNotFound).JSON(models.Resp{
				Status:  false,
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
		"product": product,
		"shop":    shop,
	})
}

func (h *OwnerHandler) Delete(c *fiber.Ctx) (err error) {
	id := new(models.IdReg)
	if err := c.QueryParser(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(id); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.handler.Delete(id); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Resp{
		Status:  true,
		Message: "success",
	})
}

func (h *OwnerHandler) Create(c *fiber.Ctx) (err error) {
	product := new(models.CreateProduct)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(500).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(product); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if value, ok := form.File["image"]; ok {
		for _, fileHeader := range value {
			product.FileName = append(product.FileName, fileHeader.Filename)
		}
	} else {
		return c.Status(500).JSON(models.Resp{
			Status:  false,
			Message: "image is missing",
		})
	}

	path, err := h.handler.Create(product)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	for index, file := range form.File["image"] {
		if err := c.SaveFile(file, "."+path.Path[index]); err != nil {
			return c.Status(500).JSON(models.Resp{
				Status:  false,
				Message: err.Error(),
			})
		}
	}

	return c.JSON(models.Resp{
		Status:  true,
		Message: "success",
	})
}

func (h *OwnerHandler) Update(c *fiber.Ctx) (err error) {
	product := new(models.CreateProduct)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(500).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(product); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if value, ok := form.File["image"]; ok {
		for _, fileHeader := range value {
			product.FileName = append(product.FileName, fileHeader.Filename)
		}
	}

	link := strings.Split(product.Link, ",")
	for _, v := range link {
		fmt.Println("product.Link->", v)
		product.FileName = append(product.FileName, v)
	}

	path, err := h.handler.Update(product)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	for _, v := range path.OldPath {
		if v != "" {
			err := os.Remove("." + v)
			if err != nil {
				return c.JSON(fiber.Map{
					"status":  500,
					"message": err.Error(),
				})
			}
		}
	}

	for index, file := range form.File["image"] {
		if err := c.SaveFile(file, "."+path.Path[index]); err != nil {
			return c.Status(500).JSON(models.Resp{
				Status:  false,
				Message: err.Error(),
			})
		}
	}

	return c.JSON(models.Resp{
		Status:  true,
		Message: "success",
	})
}

func (h *OwnerHandler) GetOrder(c *fiber.Ctx) error {
	id := new(models.IdReg)
	if err := c.QueryParser(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(id); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	products, err := h.handler.GetOrder(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
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

func (h *OwnerHandler) Login(c *fiber.Ctx) (err error) {
	param := new(models.Login)
	if err := c.BodyParser(&param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	t, id, err := h.handler.Login(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"token":   t,
		"shop_id": id,
	})
}

func (h *OwnerHandler) SaveImage(c *fiber.Ctx) error {
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
		return c.Status(500).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	path, err := h.handler.SaveImage(id, file.Filename)
	if err != nil {
		return c.Status(500).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := c.SaveFile(file, "."+path); err != nil {
		return c.Status(500).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"Message": "Image uploaded successfully",
	})
}

func (h *OwnerHandler) DeleteImage(c *fiber.Ctx) error {
	image := new(models.Image)
	if err := c.BodyParser(image); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	validate = validator.New()
	if err := validate.Struct(image); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	// delete image from ./images
	err := os.Remove(fmt.Sprintf("./images/product/%d/%s", image.Id, image.Name))
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	if err := h.handler.DeleteImage(image); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  201,
		"message": "Image deleted successfully",
	})
}

func (h *OwnerHandler) Issued(c *fiber.Ctx) error {
	param := new(models.Issued)
	if err := c.BodyParser(param); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	validate = validator.New()
	if err := validate.Struct(param); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.handler.Issued(param); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Resp{
		Status:  true,
		Message: "success",
	})
}

func (h *OwnerHandler) GetOrders(c *fiber.Ctx) error {
	param := new(models.OwnerFillter)
	if err := c.QueryParser(param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(param); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	products, err := h.handler.GetOrders(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
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

func (h *OwnerHandler) GetCatalog(c *fiber.Ctx) error {
	param := new(models.CatalogFilter)
	if err := c.QueryParser(param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(param); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	products, err := h.handler.GetCatalog(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
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

func (h *OwnerHandler) UpdateEmail(c *fiber.Ctx) error {
	param := new(models.EmailUser)
	if err := c.BodyParser(param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(param); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.handler.UpdateEmail(param); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
	})
}

func (h *OwnerHandler) GetProfile(c *fiber.Ctx) error {
	id := new(models.IdReg)
	if err := c.QueryParser(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(id); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	user, err := h.handler.GetProfile(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.Status(fiber.StatusNotFound).JSON(models.Resp{
				Status:  false,
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
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

func (h *OwnerHandler) UpdatePassword(c *fiber.Ctx) error {
	password := new(models.Password)
	if err := c.BodyParser(password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(password); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.handler.UpdatePassword(password); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
	})
}

func (h *OwnerHandler) UpdateProfile(c *fiber.Ctx) error {
	param := new(models.DTOowner)
	if err := c.BodyParser(param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(param); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := h.handler.UpdateProfile(param); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success",
	})
}

func (h *OwnerHandler) MainPage(c *fiber.Ctx) error {
	id := new(models.IdReg)
	if err := c.QueryParser(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	validate = validator.New()
	if err := validate.Struct(id); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	info, products, err := h.handler.MainPage(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":   true,
		"message":  "success",
		"stat":     info,
		"products": products,
	})
}
