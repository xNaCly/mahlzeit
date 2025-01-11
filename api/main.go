package main

import (
	"embed"
	_ "embed"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xnacly/mahlzeit/assert"
	"github.com/xnacly/mahlzeit/database"
	"github.com/xnacly/mahlzeit/models"
	"github.com/xnacly/mahlzeit/routes"
)

//go:embed dist
var WEB_FS embed.FS

func main() {
	withDefaultFlag := flag.Bool("withDefault", false, "insert default meals into database")
	flag.Parse()
	db, err := database.New(*withDefaultFlag)
	assert.NoError(err, "ctx", "Failed to get a database instance")

	app := fiber.New(fiber.Config{
		AppName:           "mahlzeit",
		ServerHeader:      "mahlzeit",
		EnablePrintRoutes: true,
		DisableKeepalive:  true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if val, ok := err.(*fiber.Error); ok {
				c.Status(val.Code).JSON(models.ApiResponse{
					Success: false,
					Message: val.Message,
				})
				return nil
			}
			return nil
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

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
	}))

	routes.RegisterRoutes(app, routes.Routes...)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(WEB_FS),
		PathPrefix: "/dist",
	}))

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
