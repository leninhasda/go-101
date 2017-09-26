package main

func s(sr string) {
	println(sr)
}

func main() {
	msgs := make(chan string, 2)
	// chn2 := make(chan string)

	msgs <- "msg 1"
	// msgs <- "msg 2"
	// msgs <- "msg 3"

	// fmt.Println(<-msgs)
	// chn2 <- msgs
	s(<-msgs) // can directly send the string buffer

	// fmt.Println(<-msgs)
	// fmt.Println(<-msgs)

}
