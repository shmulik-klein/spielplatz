package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shmulik-klein/spielplatz/coffee-project/handlers"
)

func main() {
	logger := log.New(os.Stdout, "coffee-api", log.LstdFlags)

	hh := handlers.NewHello(logger)
	coffeesHandler := handlers.NewCoffees(logger)
	mux := http.NewServeMux()
	mux.Handle("/hello", hh)
	mux.Handle("/coffees", coffeesHandler)

	server := &http.Server{
		Addr:         ":9000",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan
	logger.Printf("Recieved %s signal, gracefully terminated.\n", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
