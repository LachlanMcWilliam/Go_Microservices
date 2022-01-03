package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	// provides the capability to pass a logger to the struct using dependency injection
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// The (h *Hello) part of the function signature provides the ability to add the function onto the Hello struct
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Breaking this line down
	// h is the higher level struct that this method is attached to
	// l is the logger attached to the struct
	h.l.Println("Hello world!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", d)
}
