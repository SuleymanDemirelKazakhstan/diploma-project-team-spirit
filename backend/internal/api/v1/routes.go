package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"secondChance/internal/api/v1/handlers"
	"secondChance/internal/api/v1/middlewares"
)

func Routes(app *fiber.App, h *handlers.Handler) {
	//public service for health check service
	app.Get("/", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": true, "message": "success"}) })
	app.Get("/pgversion", h.CheckPostgresVersion)

	// Auth
	auth := app.Group("/auth")
	auth.Post("/customer", h.LoginCustomer)
	auth.Post("/owner", h.LoginOwner)

	// Admin
	v1 := app.Group("/admin")
	v1.Use(basicauth.New(basicauth.Config{Users: map[string]string{"admin": "admin"}}))
	v1.Post("/create", h.CreateOwner)
	v1.Delete("/delete", h.DeleteOwner)
	v1.Get("/get", h.GetOwner)
	v1.Get("/getall", h.GetAllOwner)
	v1.Put("/update", h.UpdateOwner)

	// Unauthorized customer
	customer := app.Group("/c")
	customer.Post("/signup", h.SignUp)
	customer.Get("/get", h.GetProduct)
	customer.Get("/allproduct", h.GetAllProduct)

	// Authorized customer
	customer.Use(middlewares.Protected())
	customer.Get("/buy", h.SoldProduct)

	// Owner
	owner := app.Group("/owner")
	owner.Use(middlewares.Protected())
	owner.Post("/create", h.CreateProduct)
	owner.Get("/delete", h.SoldProduct)
	owner.Get("/get", h.GetProduct)
	owner.Get("/getall", h.GetAllProduct)
	owner.Put("/update", h.UpdateProduct)
}
