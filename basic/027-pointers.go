package main

import "fmt"

type shape struct {
	width, height int
}

func (s *shape) area() int {
	return s.width * s.height
}

func (s shape) perim() int {
	return s.width*2 + s.height*2
}

func main() {

	s := shape{10, 2}
	fmt.Println(s.perim())
}
