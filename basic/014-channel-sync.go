package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 2)
	go worker1(done)
	go worker2(done)

	<-done
	<-done
}

func worker1(done chan bool) {
	fmt.Println("worker 1 started ...")
	time.Sleep(time.Second * 2)
	fmt.Println("worker 1 done.")
	done <- true
}

func worker2(done chan bool) {
	fmt.Println("worker 2 started ...")
	time.Sleep(time.Second)
	fmt.Println("worker 2 done.")
	done <- true
}
