package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(os.Getenv("GOPATH"))

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		k, v := pair[0], pair[1]
		fmt.Println(k, "=", v)
	}

}
