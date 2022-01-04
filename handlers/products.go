package handlers

import (
	"log"
	"net/http"

	"github.com/lachlanmcwilliam/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Test the method of the request
	switch r.Method {
	// If the request is a GET, call the GetProducts function
	case http.MethodGet:
		p.GetProducts(w, r)
	// Catch all other methods and return a 405 Method Not Allowed
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
