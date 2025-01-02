package database

import (
	"log/slog"

	"github.com/xnacly/mahlzeit/models"
)

func (d *Database) Meals() ([]models.Meal, error) {
	meals, err := d.MealsShallow()
	if err != nil {
		slog.Error("Meals", "ctx", "failed to call Database.MealsShallow", "err", err)
		return nil, err
	}
	d.m.Lock()
	for i, meal := range meals {
		ingredientsQuery, err := d.conn.Query("SELECT name, unit, amount FROM ingredients WHERE meal_id = ?", meal.Id)
		if err != nil {
			slog.Error("Meals", "ctx", "ingredientsQuery", "err", err)
			continue
		}
		for ingredientsQuery.Next() {
			ingredient := models.Ingredient{}
			err := ingredientsQuery.Scan(&ingredient.Name, &ingredient.Name, &ingredient.Unit)
			if err != nil {
				slog.Error("Meals", "ctx", "ingredientsQuery.Scan", "err", err)
				continue
			}
			meals[i].Ingredients = append(meals[i].Ingredients, ingredient)
		}
	}
	d.m.Unlock()
	return meals, nil
}

// MealsShallow returns a list of found meals, without their ingredients
func (d *Database) MealsShallow() ([]models.Meal, error) {
	d.m.Lock()
	mealsQuery, err := d.conn.Query("SELECT id, name, image, recipe FROM meals")
	if err != nil {
		slog.Error("MealsShallow", "ctx", "mealsQuery", "err", err)
		return []models.Meal{}, nil
	}

	meals := make([]models.Meal, 0, 16)
	for mealsQuery.Next() {
		m := models.Meal{}
		if err := mealsQuery.Scan(&m.Id, &m.Name, &m.Image, &m.Ingredients); err != nil {
			slog.Error("MealsShallow", "ctx", "mealsQuery.Scan", "err", err)
			continue
		}
	}
	d.m.Unlock()

	return meals, nil
}

func (d *Database) MealById(id int) (models.Meal, error) {
	empty := models.Meal{}
	mealQuery := d.conn.QueryRow("SELECT id, name, image, recipe FROM meals WHERE id = ?", id)
	meal := models.Meal{}
	if err := mealQuery.Scan(); err != nil {
		slog.Error("MealById", "ctx", "mealsQuery.Scan", "err", err)
		return empty, err
	}

	ingredientsQuery, err := d.conn.Query("SELECT name, unit, amount FROM ingredients WHERE meal_id = ?", meal.Id)
	if err != nil {
		slog.Error("MealById", "ctx", "ingredientsQuery", "err", err)
		return empty, err
	}
	for ingredientsQuery.Next() {
		ingredient := models.Ingredient{}
		err := ingredientsQuery.Scan(&ingredient.Name, &ingredient.Name, &ingredient.Unit)
		if err != nil {
			slog.Error("MealById", "ctx", "ingredientsQuery.Scan", "err", err)
			return empty, err
		}
		meal.Ingredients = append(meal.Ingredients, ingredient)
	}

	return meal, nil
}
