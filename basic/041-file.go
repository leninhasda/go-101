package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var p = fmt.Print
var pln = fmt.Println
var pf = fmt.Printf

func main() {
	// d1 := []byte("hello\ngo\n")
	// err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	// file := "/tmp/defer.txt"
	file := "/tmp/dat1"

	dat, err := ioutil.ReadFile(file)
	check(err)
	p(string(dat))

	f, err := os.Open(file)
	defer f.Close()
	check(err)

	b1 := make([]byte, 1)
	n1, err := f.Read(b1)
	check(err)
	pf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(2, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	pf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(1, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	pf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(3)
	check(err)
	pf("3 bytes: %s\n", string(b4))

	f2, err := os.Create("/tmp/dat2")
	check(err)
	defer f2.Close()
}
