package main

import (
	"errors"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xnacly/mahlzeit/database"
	"github.com/xnacly/mahlzeit/models"
	"github.com/xnacly/mahlzeit/routes"

	"log/slog"
)

func main() {
	db, err := database.New()
	if err != nil {
		slog.Error("Failed to get a database instance", "err", err)
		os.Exit(1)
	}
	database.Set(db)

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
			database.Get().Destroy()
		}
	}()
}
