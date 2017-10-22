package main

import (
	"encoding/json"
	"net/http"
)

func methodNotAllwed(w http.ResponseWriter, r *http.Request) {
	d := "method not allowed"

	res := response{}
	res.statusCode(405).data(d).serve(w)
}

func parseBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

//  simple response structure
type response struct {
	w    http.ResponseWriter
	code int
	d    interface{}
}

func (r *response) statusCode(c int) *response {
	r.code = c
	return r
}
func (r *response) data(d interface{}) *response {
	r.d = d
	return r
}

func (r *response) serve(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.code)
	json.NewEncoder(w).Encode(r.d)
}
