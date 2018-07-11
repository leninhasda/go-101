package main

import (
	"fmt"
	"net"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	ln, err := net.Listen("tcp", ":9001")
	if err != nil {
		spew.Dump(err)
	}
	// go func() {
	for {
		con, err := ln.Accept()
		if err != nil {

		}
		fmt.Println("receive")
		con.Write([]byte(`hello world!`))
		fmt.Println("sent")
		// spew.Dump("received", con)
	}
	// }()
	// time.S
	// conn, err := net.Dial("tcp", ":9001")
	// if err != nil {
	// 	spew.Dump(err)
	// }
	// // spew.Dump(conn)
	// bt := make([]byte, 12)
	// // bt := []byte{}
	// conn.Read(bt)
	// spew.Dump(string(bt))
}
