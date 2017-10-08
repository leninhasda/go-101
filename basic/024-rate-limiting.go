package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 6)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// requests <- 10 // can't do this after close

	limiter := time.Tick(time.Millisecond * 500)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			fmt.Println("inside routine")
			burstyLimiter <- t
			fmt.Println("inside asfd routine")
		}
	}()

	fmt.Println(burstyLimiter)
	time.Sleep(time.Second)
	<-burstyLimiter
	time.Sleep(time.Second * 2)

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyRequests
		fmt.Println("request", req, time.Now())
	}
}
