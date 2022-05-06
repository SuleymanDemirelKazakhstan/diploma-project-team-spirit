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

	// Auth
	auth := app.Group("/auth")
	auth.Post("/customer", h.Customer.Login)
	auth.Post("/owner", h.Shop.Login)

	// Admin
	admin := app.Group("/admin")
	admin.Use(basicauth.New(basicauth.Config{Users: map[string]string{"admin": "admin"}}))
	admin.Post("/create", h.Admin.Create)
	admin.Delete("/delete", h.Admin.Delete)
	admin.Get("/get", h.Admin.Get)
	admin.Get("/getall", h.Admin.GetAll)
	admin.Put("/update", h.Admin.Update)
	admin.Post("/saveimage", h.Admin.SaveImage)
	admin.Delete("/deleteimage", h.Admin.DeleteImage)

	// Unauthorized customer
	guest := app.Group("/g")
	guest.Post("/signup", h.Customer.SignUp)
	guest.Get("/get", h.Shop.Get)
	guest.Get("/allproduct", h.Shop.GetAll)

	// Authorized customer
	customer := app.Group("/c")
	customer.Use(middlewares.Protected())
	customer.Get("/buy", h.Customer.Buy)

	// Owner
	owner := app.Group("/owner")
	owner.Use(middlewares.Protected())
	owner.Post("/create", h.Shop.Create)
	owner.Delete("/delete", h.Shop.Delete)
	owner.Get("/get", h.Shop.Get)
	owner.Get("/getall", h.Shop.GetAll)
	owner.Put("/update", h.Shop.Update)
}
