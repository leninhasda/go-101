package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 2)

	go func() {
		time.Sleep(time.Second * 2)
		c <- "task 1"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		c <- "task 2"
	}()

	// for i := 0; i < 2; i++ {
	select {
	case msg1 := <-c:
		fmt.Println(msg1)
	case msg2 := <-c:
		fmt.Println(msg2)
	}
	// }
}
