package database

import (
	"database/sql"
	"log/slog"

	"github.com/xnacly/mahlzeit/models"
)

func (d *Database) Meals() ([]models.Meal, error) {
	meals, err := d.MealsShallow()
	if err != nil {
		slog.Error("Meals", "ctx", "failed to call Database.MealsShallow", "err", err)
		return nil, err
	}
	for i, meal := range meals {
		ingredientsQuery, err := d.conn.Query("SELECT name, unit, kind, amount FROM ingredients WHERE meal_id = ?", meal.Id)
		if err != nil {
			slog.Error("Meals", "ctx", "ingredientsQuery", "err", err)
			continue
		}
		for ingredientsQuery.Next() {
			ingredient := models.Ingredient{}
			err := ingredientsQuery.Scan(&ingredient.Name, &ingredient.Unit, &ingredient.Kind, &ingredient.Amount)
			if err != nil {
				slog.Error("Meals", "ctx", "ingredientsQuery.Scan", "err", err)
				continue
			}
			meals[i].Ingredients = append(meals[i].Ingredients, ingredient)
		}
	}
	return meals, nil
}

// MealsShallow returns a list of found meals, without their ingredients
func (d *Database) MealsShallow() ([]models.Meal, error) {
	mealsQuery, err := d.conn.Query("SELECT id, name, minutes, image, recipe FROM meals")
	if err != nil {
		slog.Error("MealsShallow", "ctx", "mealsQuery", "err", err)
		return []models.Meal{}, nil
	}

	meals := make([]models.Meal, 0, 16)
	for mealsQuery.Next() {
		m := models.Meal{}
		if err := mealsQuery.Scan(&m.Id, &m.Name, &m.Minutes, &m.Image, &m.Recipe); err != nil {
			slog.Error("MealsShallow", "ctx", "mealsQuery.Scan", "err", err)
			continue
		}
		meals = append(meals, m)
	}

	return meals, nil
}

func (d *Database) MealById(id int) (models.Meal, error) {
	empty := models.Meal{}
	mealQuery := d.conn.QueryRow("SELECT id, name, minutes, image, recipe FROM meals WHERE id = ?", id)
	meal := models.Meal{}
	if err := mealQuery.Scan(&meal.Id, &meal.Name, &meal.Minutes, &meal.Image, &meal.Recipe); err != nil {
		slog.Error("MealById", "ctx", "mealsQuery.Scan", "err", err)
		return empty, err
	}

	ingredientsQuery, err := d.conn.Query("SELECT name, unit, kind, amount FROM ingredients WHERE meal_id = ?", meal.Id)
	if err != nil {
		slog.Error("MealById", "ctx", "ingredientsQuery", "err", err)
		return empty, err
	}
	for ingredientsQuery.Next() {
		ingredient := models.Ingredient{}
		err := ingredientsQuery.Scan(&ingredient.Name, &ingredient.Unit, &ingredient.Kind, &ingredient.Amount)
		if err != nil {
			slog.Error("MealById", "ctx", "ingredientsQuery.Scan", "err", err)
			return empty, err
		}
		meal.Ingredients = append(meal.Ingredients, ingredient)
	}

	return meal, nil
}

func (d *Database) DeleteById(id int) error {
	d.m.Lock()
	defer d.m.Unlock()
	_, err := d.conn.Exec("DELETE FROM ingredients WHERE meal_id = ?", id)
	if err != nil {
		slog.Error("DeleteById", "ctx", "DELETE FROM ingredients WHERE meal_id = ?", "id", id, "err", err)
		return nil
	}
	_, err = d.conn.Exec("DELETE FROM meals WHERE id = ?")
	return err
}

func (d *Database) NewMeal(m models.Meal) error {
	d.m.Lock()
	defer d.m.Unlock()
	tx, err := d.conn.Begin()
	if err != nil {
		slog.Error("NewMeal", "ctx", "start transaction", "err", err)
		return err
	}

	err = func(transaction *sql.Tx) error {
		r, err := d.conn.Exec("INSERT INTO meals (name, minutes, image, recipe) VALUES(?, ?, ?, ?)", m.Name, m.Minutes, m.Image, m.Recipe)
		if err != nil {
			slog.Error("NewMeal", "ctx", "INSERT INTO meals (name, minutes, image, recipe) VALUES(?, ?, ?, ?)", "Meal", m, "err", err)
			return err
		}

		lastId, err := r.LastInsertId()
		if err != nil {
			return err
		}

		for _, ingredient := range m.Ingredients {
			if err := d.NewIngredient(ingredient, lastId); err != nil {
				return err
			}
		}
		return nil
	}(tx)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (d *Database) NewIngredient(ingredient models.Ingredient, mealId int64) error {
	d.m.Lock()
	defer d.m.Unlock()
	_, err := d.conn.Exec("INSERT INTO ingredients (meal_id, name, unit, kind, amount) VALUES(?, ?, ?, ?)", mealId, ingredient.Name, ingredient.Unit, ingredient.Kind, ingredient.Amount)
	return err
}

func (d *Database) Ingredients() ([]models.Ingredient, error) {
	ingredientsQuery, err := d.conn.Query("SELECT name, unit, kind, amount FROM ingredients")
	ingredients := []models.Ingredient{}
	if err != nil {
		slog.Error("Ingredients", "ctx", "SELECT name, unit, kind, amount FROM ingredients", "err", err)
		return []models.Ingredient{}, err
	}
	for ingredientsQuery.Next() {
		ingredient := models.Ingredient{}
		err := ingredientsQuery.Scan(&ingredient.Name, &ingredient.Unit, &ingredient.Kind, &ingredient.Amount)
		if err != nil {
			slog.Error("Ingredients", "ctx", "ingredientsQuery.Scan", "err", err)
			return []models.Ingredient{}, err
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}
