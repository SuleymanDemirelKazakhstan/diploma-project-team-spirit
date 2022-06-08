package api

import (
	"secondChance/internal/api/v1/handlers"
	"secondChance/internal/api/v1/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
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
	guest.Get("/verify", h.Customer.GmailCode)
	guest.Get("/discountproduct", h.Customer.GetDiscountProducts)
	guest.Get("/allshop", h.Admin.GetAll)
	guest.Get("/search", h.Customer.Search)
	guest.Get("/filter", h.Customer.GetFilter)

	// Authorized customer
	customer := app.Group("/c")
	customer.Use(middlewares.Protected())
	customer.Post("/buy", h.Customer.Buy)
	customer.Get("/getter", h.Customer.Getter)
	customer.Post("/setter", h.Customer.Setter)
	customer.Post("/saveimage", h.Customer.SaveImage)
	customer.Delete("/deleteimage", h.Customer.DeleteImage)
	customer.Get("/order", h.Customer.GetOrder)
	customer.Get("/verify", h.Customer.GmailCode)
	customer.Get("/myproducts", h.Customer.GetAllMyProduct)
	customer.Post("/updatepassword", h.Customer.UpdatePassword)

	// Owner
	owner := app.Group("/owner")
	owner.Use(middlewares.Protected())
	owner.Post("/create", h.Shop.Create)
	owner.Delete("/delete", h.Shop.Delete)
	owner.Get("/get", h.Shop.Get)
	owner.Get("/getall", h.Shop.GetAll)
	owner.Put("/update", h.Shop.Update)
	owner.Post("/saveimage", h.Shop.SaveImage)
	owner.Delete("/deleteimage", h.Shop.DeleteImage)
	owner.Post("/saveicon", h.Admin.SaveImage)
	owner.Delete("/deleteicon", h.Admin.DeleteImage)
	owner.Get("/order", h.Shop.GetOrder)
	owner.Put("/issued", h.Shop.Issued)
	owner.Get("/verify", h.Customer.GmailCode)
	owner.Put("/updateinfo", h.Admin.Update)
	owner.Get("/myproducts", h.Shop.GetAllMyProduct)
}
