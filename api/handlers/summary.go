package handlers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/xnacly/mahlzeit/database"
	"github.com/xnacly/mahlzeit/models"
)

type summary struct {
	IngredientsByKind map[models.Kind][]models.Ingredient `json:"ingredients_by_kind"`
}

func Summary(c *fiber.Ctx) error {
	ids := c.Query("meal_ids")
	if ids == "" {
		return fiber.ErrBadRequest
	}
	db := c.Locals("db").(*database.Database)
	s := summary{
		IngredientsByKind: map[models.Kind][]models.Ingredient{},
	}
	for _, rawId := range strings.Split(ids, ",") {
		id, err := strconv.ParseInt(rawId, 10, 32)
		if err != nil {
			return fiber.ErrBadRequest
		}

		meal, err := db.MealById(int(id))
		if err != nil {
			return fiber.ErrNotFound
		}
		for _, i := range meal.Ingredients {
			s.IngredientsByKind[i.Kind] = append(s.IngredientsByKind[i.Kind], i)
		}
	}
	return c.JSON(models.ApiResponse{
		Success: true,
		Code:    200,
		Message: "pong",
		Data:    s,
	})
}
