package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"secondChance/internal/models"
)

var validate *validator.Validate

func (h *Handler) CreateOwner(c *fiber.Ctx) (err error) {
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

	hash, err := HashPassword(a.Password)
	if err!= nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	a.Password = hash
	if err := h.ServiceLayer.CreateOwner(a); err != nil {
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

func (h *Handler) DeleteOwner(c *fiber.Ctx) (err error) {
	a := new(models.OwnerEmailRequest)
	if err := c.QueryParser(a); err != nil {
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

	if err := h.ServiceLayer.DeleteOwner(a); err != nil {
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

func (h *Handler) GetOwner(c *fiber.Ctx) (err error) {
	a := new(models.OwnerEmailRequest)
	if err := c.QueryParser(a); err != nil {
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

	user, err := h.ServiceLayer.GetOwner(a)
	if err != nil {
		if err.Error() == "sql: no rows in result set"{
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

func (h *Handler) GetAllOwner(c *fiber.Ctx) (err error) {
	users, err := h.ServiceLayer.GetAllOwner()
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

func (h *Handler) UpdateOwner(c *fiber.Ctx) (err error) {
	email := new(models.OwnerEmailRequest)
	if err := c.QueryParser(email); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	validate = validator.New()
	if err := validate.Struct(email); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}

	userDB, err := h.ServiceLayer.GetOwner(email)
	if err != nil {
		if err.Error() == "sql: no rows in result set"{
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

	userReq := new(models.Owner)
	if err := c.BodyParser(userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResp{
			Status:  false,
			Message: err.Error(),
		})
	}

	//Todo Status Code
	user := newUser(userReq, &models.Owner{
		Email: userDB.Email,
		Name: userDB.Name,
		Phone: userDB.Phone,
		Password: userDB.Password,
	})
	if err := h.ServiceLayer.UpdateOwner(email,user); err != nil {
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

//Todo validate
func newUser(userReq, user *models.Owner) *models.Owner {
	if userReq.Email != "" {
		user.Email = userReq.Email
	}
	if userReq.Name != "" {
		user.Name = userReq.Name
	}
	if userReq.Phone != 0 {
		user.Phone = userReq.Phone
	}
	if userReq.Password != ""{
		var err error
		user.Password, err = HashPassword(userReq.Password)
		if err != nil{
			log.Fatal(err)
		}
	}
	return user
}

