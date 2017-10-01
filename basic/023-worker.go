package main

import (
	"fmt"
	"time"
)

// workers
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// run 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	// close the job channel
	close(jobs)

	// get results
	for a := 1; a <= 5; a++ {
		<-results
	}
}
