package main

import "fmt"

func main() {
	// done := make(chan string)
	// go func() { done <- "true" }()

	done := make(chan string, 1)
	// done := make(chan string)
	done <- "true"
	done <- "true"

	s := <-done
	fmt.Println("hello", s)
}
