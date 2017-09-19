package main

import "fmt"

func main() {
	// infinite style
	for {
		fmt.Println("this is a infinite loop!")
		fmt.Println("or not :p")
		break
	}

	// while style
	count := 2
	for count > 0 {
		fmt.Println(count)
		count = count - 1
	}

	// good old classic style
	for cnt := 3; cnt > 0; cnt = cnt - 1 {
		fmt.Println("cnt is ", cnt)
	}
}
