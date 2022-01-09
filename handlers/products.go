package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	// If the request is a POST, call the AddProduct function
	case http.MethodPost:
		p.AddProduct(w, r)
	// If the request is a PUT, call the UpdateProduct function, and expect the id in the URL
	case http.MethodPut:
		p.UpdateProduct(w, r)
	// Catch all other methods and return a 405 Method Not Allowed
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Endpoint Hit: GetProducts")
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Endpoint Hit: AddProduct")

	// Decode the incoming json
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)
	w.WriteHeader(http.StatusCreated)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// As this is a PUT request, we need to read the URL of the request to get the ID of the product
	id_regex := regexp.MustCompile(`/([0-9]+)`)
	g := id_regex.FindAllStringSubmatch(r.URL.Path, -1)

	if len(g) != 1 {
		p.l.Println("Error: The URI contains more than one ID")
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}

	if len(g[0]) != 2 {
		p.l.Println("Error: The URI does not contain an ID")
		p.l.Println(len(g[0]))
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}

	idString := g[0][1]
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(w, "Unable to update product", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Unable to update product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
