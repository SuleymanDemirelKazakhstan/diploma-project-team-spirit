package app

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"secondChance/internal/api/v1"
	"secondChance/internal/api/v1/handlers"
	"secondChance/internal/db"
	"secondChance/internal/models"
	"secondChance/internal/services"
	"syscall"
	"time"
)

const idleTimeout = 10 * time.Second

// Run initializes whole application
func Run() {
	_app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			//Status code default 500
			code := fiber.StatusInternalServerError

			return c.Status(code).JSON(models.ErrorResp{
				Status:  false,
				Message: err.Error(),
			})
		},
		DisableStartupMessage: true,
	})

	_db := db.NewDB()
	defer _db.Close()

	_dbm := db.NewDataBaseLayers(_db)
	_service := services.NewServiceLayer(_dbm)
	_handler := handlers.NewHandler(_service)

	api.Routes(_app, _handler)

	// Listen from a different goroutine
	go func() {
		if err := _app.Listen(":" + "8080"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create chanel to segnify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c //This block the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = _app.Shutdown()

	log.Println("Running cleanup tasks...")

}
