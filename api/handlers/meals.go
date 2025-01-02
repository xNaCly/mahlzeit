package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/xnacly/mahlzeit/database"
	"github.com/xnacly/mahlzeit/models"
)

func Meals(c *fiber.Ctx) error {
	db := c.Locals("db").(*database.Database)
	meals, err := db.Meals()
	if err != nil {
		slog.Error("mealById", "ctx", "failed to call Database.Meals", "err", err)
		return err
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Code:    200,
		Message: "Got meals",
		Data: fiber.Map{
			"meals": meals,
		},
	})
}

func MealById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		slog.Error("mealById", "ctx", "bad parameter", "err", err)
		return fiber.ErrBadRequest
	}
	db := c.Locals("db").(*database.Database)
	meals, err := db.MealById(id)
	if err != nil {
		return err
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Code:    200,
		Message: "Got meal with id: " + strconv.FormatInt(int64(id), 10),
		Data: fiber.Map{
			"meal": meals,
		},
	})
}
