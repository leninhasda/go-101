package main

import (
	"fmt"
	"net/http"
)

type customHandle struct{}

func (fn customHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a := "hello test basic auth"
	fmt.Fprint(w, a)
}

func main() {
	h := customHandle{}
	http.Handle("/", BasicAuth(h))
	http.ListenAndServe(":4000", nil)
}

// BasicAuth invokes a basic http auth
func BasicAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if ok && validate(u, p) {
			h.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="enter user / pass"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
	})
}

func validate(u, p string) bool {
	if u == "user" && p == "pass" {
		return true
	}
	return false
}
