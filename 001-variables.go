package main

import "fmt"

// there are three way of declaring variables in go

func main() {
	// 1 - declare type first assign later
	var myAge int
	var name string
	myAge = 10
	name = "faruk"
	fmt.Printf("hello my name is %s and i am %d years old\n", name, myAge)

	// 2 - declare & assign value as well as type
	var friend = "amit"
	var owe, due = 100.50, 45.00
	fmt.Printf("you are %s my friend. you owe me %.2f and due is %.2f\n", friend, owe, due)

	// 3 - directly assign value & type
	enemy := "sohel"
	willAttack := false
	if willAttack {
		fmt.Printf("enemy %s is going to attack\n", enemy)
	} else {
		fmt.Printf("enemey %s is not going to attack today\n", enemy)
	}

}
