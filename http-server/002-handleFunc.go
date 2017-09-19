package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/not-found", notFound)
	http.ListenAndServe(":4000", nil)
}

// Index handler for / resource
func Index(w http.ResponseWriter, r *http.Request) {
	a := "hello from index router"
	fmt.Fprintln(w, a)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
