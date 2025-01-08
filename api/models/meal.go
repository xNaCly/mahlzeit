package models

import "time"

type Ingredient struct {
	Name   string `json:"name"`
	Unit   string `json:"unit"`
	Kind   Kind   `json:"kind"`
	Amount int    `json:"amount"`
}

type Meal struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	Ingredients []Ingredient  `json:"ingredients,omitempty"`
	Minutes     time.Duration `json:"minutes"`
	// Cost is computed based on the minutes it takes to cook the meal, the
	// amount of ingredients, their amounts and their weight according to their
	// kind, cost is not estimated in a currency, it is simply meant to compare
	// meals in mahlzeit
	Cost   int    `json:"cost"`
	Image  string `json:"image"`  // image url, maybe stockfoto integration?
	Recipe string `json:"recipe"` // typically an url
}
