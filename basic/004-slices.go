package main

import "fmt"

func main() {
	intSlice := []int{}

	fmt.Println(len(intSlice))

	intSlice = append(intSlice, 10)
	intSlice = append(intSlice, 0)
	intSlice = append(intSlice, 2)

	fmt.Println(len(intSlice))
	fmt.Println(intSlice)

	boolType := make([]bool, 2)
	boolType[0] = true
	boolType = append(boolType, !false)
	fmt.Println(len(boolType))
	fmt.Println(boolType)

}
