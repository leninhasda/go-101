package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var p = fmt.Println
var pf = fmt.Printf

func main() {

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat3", d1, 0644)
	check(err)
	defer os.Remove("/tmp/dat3")

	f, err := os.Create("/tmp/dat4")
	check(err)
	defer f.Close()
	defer os.Remove("/tmp/dat4")

	d2 := []byte{115, 111, 100, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	pf("wrote %d bytest\n", n2)

	n3, err := f.WriteString("writes\n")
	pf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	pf("wrote %d bytes\n", n4)

	w.Flush()
}
