package main

import (
	"errors"
	"fmt"
)

func errorInLast() (int, bool, string, error) {
	return 1, false, "name", errors.New("error is the last param")
}

type customError struct {
	code int
	msg  string
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s\n", e.code, e.msg)
}

func usingCustromErr() error {
	return &customError{code: 101, msg: "yoo error!"}
}

func main() {
	i, b, s, r := errorInLast()
	fmt.Println(i, b, s, r)

	ce := usingCustromErr()
	fmt.Println(ce)
}
