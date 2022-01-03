package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lachlanmcwilliam/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// Create a new multiplexer to handle the routes
	sm := http.NewServeMux()
	// Register the handler for the route "/"
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// Create a new server to listen on port 9090 and pass in the multiplexer
	http.ListenAndServe(":9090", sm)
}
