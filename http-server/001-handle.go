package main

import (
	"fmt"
	"net/http"
)

type helloHandle struct{}

func (fn helloHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a := "hello handle"
	fmt.Fprint(w, a)
}

func main() {
	h := helloHandle{}
	http.Handle("/", h)
	http.ListenAndServe(":4000", nil)
}
