package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Create a channel to listen for an interrupt or terminate signal from the OS.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until we receive our signal.
	sig := <-sigChan
	l.Println("Received shutdown signal, gracefully shutting down:", sig)

	// Create context with timeout
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// Shutdown the server and clean up resources
	s.Shutdown(ctx)
}
