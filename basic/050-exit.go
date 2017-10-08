package main

import "os"

func main() {
	// this will not run !
	defer println("!!!")

	os.Exit(500)
}
