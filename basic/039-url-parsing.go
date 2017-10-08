package main

import (
	"fmt"
	"net/url"
)

var p = fmt.Println

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	p(u.Scheme)
	p(u.User)
	p(u.User.Username())
	ps, _ := u.User.Password()
	p(ps)

	p(u.Host)
	p(u.Hostname())
	p(u.Port())

	p(u.Path)
	p(u.Fragment)

	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"][0])
}
