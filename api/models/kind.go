package models

type Kind string

const (
	KIND_OILS_N_VINEGAR Kind = "Oils & Vinegar"
	KIND_MEAT_N_FISH    Kind = "Meat & Fish"
	KIND_SPICES         Kind = "Spices"
	KIND_DIARY          Kind = "Diary"
	KIND_CARBS          Kind = "Carbs"
	KIND_VEGETABLES     Kind = "Vegetables"
	KIND_SAUCES         Kind = "Sauces"
)

func truncateToHundred(i int) int {

	return i
}

func Cost(m Meal) int {
	weights := map[Kind]int{
		KIND_OILS_N_VINEGAR: 7,
		KIND_MEAT_N_FISH:    6,
		KIND_SPICES:         5,
		KIND_DIARY:          4,
		KIND_CARBS:          3,
		KIND_VEGETABLES:     2,
		KIND_SAUCES:         1,
	}

	cost := len(m.Ingredients) / int(m.Minutes)
	for _, i := range m.Ingredients {
		cost += weights[i.Kind]
	}

	cost = cost % 100
	if cost == 0 {
		cost = 100
	}

	return cost
}
