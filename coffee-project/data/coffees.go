package data

import (
	"encoding/json"
	"io"
	"time"
)

type CoffeeCup struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	DrinkedOn string `json:"drinked_on"`
}

type CoffeeCups []*CoffeeCup

func (c *CoffeeCups) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(c)
}

func GetCoffees() CoffeeCups {
	return coffeeList
}

var coffeeList = CoffeeCups{
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
