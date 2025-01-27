package handlers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/xnacly/mahlzeit/database"
	"github.com/xnacly/mahlzeit/models"
)

type summary struct {
	Meals             []string
	IngredientsByKind map[models.Kind][]models.Ingredient `json:"ingredients_by_kind"`
}

func Summary(c *fiber.Ctx) error {
	ids := c.Query("meal_ids")
	if ids == "" {
		return fiber.ErrBadRequest
	}
	db := c.Locals("db").(*database.Database)
	s := summary{
		Meals:             []string{},
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
		s.Meals = append(s.Meals, meal.Name)
		for _, i := range meal.Ingredients {
			s.IngredientsByKind[i.Kind] = append(s.IngredientsByKind[i.Kind], i)
		}
	}
	if c.QueryBool("text", false) {
		b := strings.Builder{}
		b.WriteString("## Meals:\n")
		for i, meal := range s.Meals {
			b.WriteString(meal)
			if i != len(s.Meals) {
				b.WriteRune('\n')
			}
		}
		b.WriteString("## Ingredients:\n")
		for group, ingredients := range s.IngredientsByKind {
			b.WriteString("### ")
			b.WriteString(string(group))
			b.WriteRune('\n')
			for i, ingredient := range ingredients {
				b.WriteString("- [ ] ")
				b.WriteString(ingredient.Name)
				b.WriteString(" (")
				b.WriteString(strconv.FormatInt(int64(ingredient.Amount), 10))
				b.WriteString(ingredient.Unit)
				b.WriteRune(')')

				if i != len(ingredients) {
					b.WriteRune('\n')
				}
			}
		}
		_, err := c.Status(fiber.StatusOK).WriteString(b.String())
		return err
	} else {
		return c.JSON(models.ApiResponse{
			Success: true,
			Code:    200,
			Message: "pong",
			Data:    s,
		})
	}
}
