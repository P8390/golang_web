package main

import (
	"fmt"
	"net/http" // All the HTTP calls are handled by this library
)

func main() {
	// This is request handler which invoke handler func for given url pattern
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World with URL - %s", r.URL.Path) // This will go as output
	})
	fs := http.FileServer(http.Dir("/")) // if folder is having index.html, that will be returned default else list of files

	http.Handle("/", fs)

	// This listen default on port 80 and serve the request
	http.ListenAndServe(":9090", nil)
}
