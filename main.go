package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lachlanmcwilliam/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// Create a new multiplexer to handle the routes
	sm := http.NewServeMux()

	// Register the handlers
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// Create a new server passing the multiplexer
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
