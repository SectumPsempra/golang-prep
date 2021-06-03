package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// HandleFunc is a convenience method on go http package
	// it registers the function and the path on the single default servemux(http handler)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
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
	})

	http.HandleFunc("/abc", func(http.ResponseWriter, *http.Request) {
		log.Println("Print abc, not hello world!!")
	})

	// constructs an http server and registers a default handler to it, second param
	// is handler(if nil, will use default servemux)->
	http.ListenAndServe(":9090", nil)
}
