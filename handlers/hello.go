package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// return Hello handler as a reference
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	// read the request
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// WriteHeader lets us specify the status code
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("OOps"))

		// above two lines are equivalent to writing below code:
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	log.Printf("Data %s\n", d)

	// print the response
	fmt.Fprintf(rw, "Hello %s\n", d)
}
