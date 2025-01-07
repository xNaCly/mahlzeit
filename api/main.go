package main

import (
	"errors"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xnacly/mahlzeit/assert"
	"github.com/xnacly/mahlzeit/database"
	"github.com/xnacly/mahlzeit/models"
	"github.com/xnacly/mahlzeit/routes"
)

func main() {
	withDefaultFlag := flag.Bool("withDefault", false, "insert default meals into database")
	flag.Parse()
	db, err := database.New(*withDefaultFlag)
	assert.NoError(err, "ctx", "Failed to get a database instance")

	app := fiber.New(fiber.Config{
		AppName:           "mahlzeit",
		ServerHeader:      "mahlzeit",
		EnablePrintRoutes: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return c.Status(code).JSON(models.ApiResponse{
				Success: false,
				Code:    code,
				Message: err.Error(),
				Data:    nil,
			})
		},
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowMethods:     "",
		Next:             nil,
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
	}))

	routes.RegisterRoutes(app, routes.Routes...)

	app.Use(func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Requested ressource not found")
	})

	log.Fatal(app.Listen(":8080"))
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for range c {
			slog.Info("destroying database connection...")
			assert.NoError(db.Destroy(), "ctx", "failed to destroy db connection")
		}
	}()
}
