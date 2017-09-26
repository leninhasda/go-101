package main

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
	// fmt.Println(<-pongs) // can't do this
	// fmt.Println(<-pings) // can do this
	// pings <- "msg" // can't do this
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	// fmt.Println(<-pongs)
}
