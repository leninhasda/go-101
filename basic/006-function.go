package main

import "fmt"

func main() {
	sum(10, 30)
	fmt.Println("catching values returned by functions are not mendatory!")

	add(0, 1.0)
}

func sum(a, b int) (int, bool) {
	c := a + b
	bl := true

	return c, bl

}

func add(a, b int) (c int, bl bool) {
	c = a + b
	bl = false
	fmt.Println("here is another way to declare & return values from function")

	return
}
