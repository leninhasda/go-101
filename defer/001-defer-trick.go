package main

import "fmt"

func main() {
	str := "\nPlease reverse me!"
	for _, c := range str {
		defer fmt.Printf("%s", string(c))
	}
}
