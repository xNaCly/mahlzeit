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
		Path:    "/api/ping/",
		Handler: handlers.Ping,
	},
	{
		Path:    "/api/meals/",
		Handler: handlers.Meals,
	},
	{
		Path:    "/api/meals/random",
		Handler: handlers.RandMeals,
	},
	{
		Path:    "/api/meals/summary",
		Handler: handlers.Summary,
	},
	{
		Path:    "/api/meals/:id",
		Handler: handlers.MealById,
	},
	{
		Path:    "/api/meals/",
		Handler: handlers.NewMeal,
		Method:  fiber.MethodPost,
	},
}

func RegisterRoutes(app *fiber.App, routes ...Route) {
	for _, route := range routes {
		if len(route.Method) == 0 {
			route.Method = "GET"
		}
		app.Add(route.Method, route.Path, route.Handler)
	}
}
