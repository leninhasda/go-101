package main

import "os"

func main() {

	panic("a problem") // panics here and exit

	_, err := os.Create("/tmp/fileln")
	if err != nil {
		panic(err)
	}
}
