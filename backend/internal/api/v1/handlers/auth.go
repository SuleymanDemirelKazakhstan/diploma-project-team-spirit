package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"secondChance/internal/models"
	"time"
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

	customer, err:= h.ServiceLayer.GetCustomer(param.Email)
	if err!=nil{
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

	if !CheckPasswordHash(param.Password, customer.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"message": "Invalid email or password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = customer.Name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"token": t,
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

	owner, err:= h.ServiceLayer.GetOwner(&models.OwnerEmailRequest{Email: param.Email})
	if err!=nil{
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

	if !CheckPasswordHash(param.Password, owner.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"message": "Invalid email or password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = owner.Name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"token": t,
		"user_id": owner.Id,
	})
}