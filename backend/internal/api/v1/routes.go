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
	admin := app.Group("/admin")
	admin.Use(basicauth.New(basicauth.Config{Users: map[string]string{"admin": "admin"}}))
	admin.Post("/create", h.CreateOwner)
	admin.Delete("/delete", h.DeleteOwner)
	admin.Get("/get", h.GetOwner)
	admin.Get("/getall", h.GetAllOwner)
	admin.Put("/update", h.UpdateOwner)

	// Unauthorized customer
	guest := app.Group("/g")
	guest.Post("/signup", h.SignUp)
	guest.Get("/get", h.GetProduct)
	guest.Get("/allproduct", h.GetAllProduct)
	guest.Get("/getallauction", h.GetProductAuction)

	// Authorized customer
	customer := app.Group("/c")
	customer.Use(middlewares.Protected())
	customer.Get("/buy", h.BuyProduct)

	// Owner
	owner := app.Group("/owner")
	owner.Use(middlewares.Protected())
	owner.Post("/create", h.CreateProduct)
	owner.Delete("/delete", h.DeleteProduct) // if selled_at is empty
	owner.Get("/get", h.GetProduct)
	owner.Get("/getall", h.GetAllProduct)
	owner.Put("/update", h.UpdateProduct)
}
