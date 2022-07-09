package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shmulik-klein/spielplatz/coffee-project/handlers"
)

func main() {
	logger := log.New(os.Stdout, "coffee-api", log.LstdFlags)
	hh := handlers.NewHello(logger)
	mux := http.NewServeMux()
	mux.Handle("/hello", hh)
	http.ListenAndServe(":9000", mux)
}