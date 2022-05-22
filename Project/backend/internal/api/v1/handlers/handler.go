package handlers

import (
	"secondChance/internal/services"

	"github.com/gofiber/fiber/v2"
)

type Admin interface {
	Create(c *fiber.Ctx) (err error)
	Delete(c *fiber.Ctx) (err error)
	Get(c *fiber.Ctx) (err error)
	GetAll(c *fiber.Ctx) (err error)
	Update(c *fiber.Ctx) (err error)
	Login(c *fiber.Ctx) (err error)
	SaveImage(c *fiber.Ctx) error
	DeleteImage(c *fiber.Ctx) error
}

type Customer interface {
	SignUp(c *fiber.Ctx) error
	GetOrder(c *fiber.Ctx) error
	Login(c *fiber.Ctx) (err error)
	Buy(c *fiber.Ctx) error
	SaveImage(c *fiber.Ctx) error
	DeleteImage(c *fiber.Ctx) error
	GmailCode(c *fiber.Ctx) error
	Setter(c *fiber.Ctx) error
	Getter(c *fiber.Ctx) error
}

type Shop interface {
	GetAll(c *fiber.Ctx) (err error)
	Get(c *fiber.Ctx) (err error)
	Delete(c *fiber.Ctx) (err error)
	Create(c *fiber.Ctx) (err error)
	Update(c *fiber.Ctx) (err error)
	GetOrder(c *fiber.Ctx) error
	Login(c *fiber.Ctx) (err error)
	SaveImage(c *fiber.Ctx) error
	DeleteImage(c *fiber.Ctx) error
	Issued(c *fiber.Ctx) error
}

type Handler struct {
	Admin
	Customer
	Shop
}

func NewHandler(serviceLayer *services.Service) *Handler {
	return &Handler{
		Admin:    NewAdminHandler(serviceLayer.Admin),
		Customer: NewCustomerHandler(serviceLayer.Customer),
		Shop:     NewOwnerHandler(serviceLayer.Shop),
	}
}
