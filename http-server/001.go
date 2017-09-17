package main

import (
	"fmt"
	"net/http"
)

type lnHandler struct {
	Routes []string
}

func (lh *lnHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a := "ServeHttp method done"
	fmt.Fprint(w, a)
}

func main() {
	l := lnHandler{
		Routes: nil,
	}
	http.HandleFunc("/", methodHandle)
	http.Handle("/test", l)
	http.ListenAndServe(":1234", nil)
}
func customHandle(w http.ResponseWriter, r *http.Request) {

}
func methodHandle(w http.ResponseWriter, r *http.Request) {
	a := "hello"
	fmt.Fprint(w, a)
}
