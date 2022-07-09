package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shmulik-klein/spielplatz/coffee-project/handlers"
)

func main() {
	logger := log.New(os.Stdout, "coffee-api", log.LstdFlags)

	hh := handlers.NewHello(logger)

	mux := http.NewServeMux()
	mux.Handle("/hello", hh)

	server := &http.Server{
		Addr:         "9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server.ListenAndServe()
}
