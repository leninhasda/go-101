package main

import "fmt"

var done = make(chan bool)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
	done <- true
}

func main() {

	go f("goroutine")

	<-done

	c := make(chan struct{}) // We don't need any data to be passed, so use an empty struct
	for i := 0; i < 100; i++ {
		go func(i int) {
			// doSomething()
			if i == 99 {
				fmt.Println(i)
			} else {
				fmt.Print(i, " ")
			}
			c <- struct{}{} // signal that the routine has completed
		}(i)
	}
	fmt.Println("")

	// Since we spawned 100 routines, receive 100 messages.
	for i := 0; i < 100; i++ {
		<-c
	}
}
