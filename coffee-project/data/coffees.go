package data

import "time"

type CoffeeCup struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	DrinkedOn string `json:"drinked_on"`
}

func GetCoffees() []*CoffeeCup {
	return coffeeList
}

var coffeeList = []*CoffeeCup{
	{
		ID:        1,
		Name:      "first_cup",
		DrinkedOn: time.Now().UTC().String(),
	},
	{
		ID:        2,
		Name:      "second_cup",
		DrinkedOn: time.Now().UTC().String(),
	},
}
