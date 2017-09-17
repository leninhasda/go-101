package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", Index)
	http.ListenAndServe(":5432", nil)
}

// Index handler for / resource
func Index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
		Msg  string
	}{
		Name: "Hello",
		Msg:  "World",
	}

	render(w, "index-view.html", data)
}
