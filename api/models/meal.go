package models

type Meal struct {
	Id          int
	Name        string
	Ingredients []Ingredient
	Image       string // image url, maybe stockfoto integration?
	Recipe      string // typically an url
}

type Ingredient struct {
	Name   string
	Unit   string
	Amount int
}
