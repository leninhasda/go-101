package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int, 2)

	go worker([]string{"task 1", "task 2"}, done)

	go worker([]string{"task a", "task b", "task 3"}, done)

	tc1 := <-done
	tc2 := <-done

	fmt.Printf("%d task(s) complete!\n", tc1+tc2)
}

func worker(tasks []string, done chan int) {
	for _, task := range tasks {
		fmt.Println(task)
		time.Sleep(time.Second / 2)
	}
	done <- len(tasks)
}
