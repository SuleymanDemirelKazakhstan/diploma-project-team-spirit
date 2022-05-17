package app

import (
	"log"
	"os"
	"os/signal"
	"secondChance/internal/api/v1"
	"secondChance/internal/api/v1/handlers"
	"secondChance/internal/db"
	"secondChance/internal/services"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

const idleTimeout = 10 * time.Second

// Run initializes whole application
func Run() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	_port := os.Getenv("port")

	log.Println("2Chance Start")
	_app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			//Status code default 500
			code := fiber.StatusInternalServerError

			return c.Status(code).JSON(fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		},
		DisableStartupMessage: true,
	})
	_app.Use(cors.New(), recover.New(), logger.New())
	_app.Static("/images", "./images")

	_db := db.NewDB()
	_rdb := db.NewRedis()
	defer _db.Close()
	defer _rdb.Close()

	_dbm := db.NewDataBaseLayers(_db, _rdb)
	_service := services.NewService(_dbm, _rdb)
	_handler := handlers.NewHandler(_service)

	api.Routes(_app, _handler)

	// Listen from a different goroutine
	go func() {
		if err := _app.Listen(":" + _port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create chanel to segnify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c //This block the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = _app.Shutdown()

	log.Println("Service successful shutdown.")
}
