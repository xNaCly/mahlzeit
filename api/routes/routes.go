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
		Handler: handlers.Ping,
	},
	{
		Path:    "/meals/",
		Handler: handlers.Meals,
	},
	{
		Path:    "/meals/random",
		Handler: handlers.RandMeals,
	},
	{
		Path:    "/meals/:id",
		Handler: handlers.MealById,
	},
	{
		Path:    "/meals/",
		Handler: handlers.NewMeal,
		Method:  fiber.MethodPost,
	},
	{
		Path: "/",
		Handler: func(c *fiber.Ctx) error {
			return fiber.ErrNotFound
		},
	},
}

func RegisterRoutes(app *fiber.App, routes ...Route) {
	apiGroup := app.Group("/api")
	for _, route := range routes {
		if len(route.Method) == 0 {
			route.Method = "GET"
		}
		apiGroup.Add(route.Method, route.Path, route.Handler)
	}
}
