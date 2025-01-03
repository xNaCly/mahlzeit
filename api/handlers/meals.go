package handlers

import (
	"log/slog"
	"maps"
	"math/rand/v2"
	"slices"
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

func RandMeals(c *fiber.Ctx) error {
	db := c.Locals("db").(*database.Database)
	meals, err := db.MealsShallow()
	if err != nil {
		return err
	}

	if len(meals) < 7 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse{
			Success: true,
			Code:    fiber.StatusBadRequest,
			Message: "At least 7 meals are required to generate a weeks worth of meals",
		})
	}

	mealIdToMeal := map[int]models.Meal{}

	for len(mealIdToMeal) <= 7 {
		randMeal := meals[rand.IntN(7)]
		mealIdToMeal[randMeal.Id] = randMeal
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Code:    200,
		Message: "Got random meals",
		Data: fiber.Map{
			"meals": slices.Collect(maps.Values(mealIdToMeal)),
		},
	})
}
