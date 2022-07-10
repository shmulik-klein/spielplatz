package handlers

import (
	"log"
	"net/http"

	"github.com/shmulik-klein/spielplatz/coffee-project/data"
)

type Coffees struct {
	logger *log.Logger
}

func NewCoffees(logger *log.Logger) *Coffees {
	return &Coffees{logger}
}

func (coffees *Coffees) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	coffeesList := data.GetCoffees()
	err := coffeesList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal response.", http.StatusInternalServerError)
	}
}
