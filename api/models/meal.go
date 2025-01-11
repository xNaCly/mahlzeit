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
	Image       string        `json:"image"`  // image url, maybe stockfoto integration?
	Recipe      string        `json:"recipe"` // typically an url
}
