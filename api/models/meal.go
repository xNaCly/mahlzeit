package models

type Meal struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients,omitempty"`
	Image       string       `json:"image"`  // image url, maybe stockfoto integration?
	Recipe      string       `json:"recipe"` // typically an url
}

type Ingredient struct {
	Name   string `json:"name"`
	Unit   string `json:"unit"`
	Amount int    `json:"amount"`
}
