package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xnacly/mahlzeit/handlers"
)

type Route struct {
	Path    string
	Method  string
	Handler func(*fiber.Ctx) error
}

var Routes = []Route{
	{
		Path:    "/ping/",
		Method:  "GET",
		Handler: handlers.Ping,
	},
}

func RegisterRoutes(app *fiber.App, routes ...Route) {
	for _, route := range routes {
		app.Add(route.Method, route.Path, route.Handler)
	}
}
