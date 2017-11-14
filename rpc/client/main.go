package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	serverAddress := "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{2, 5}
	var reply int
	er := client.Call("Arith.Multiply", args, &reply)
	if er != nil {
		log.Fatal("arith error:", er)
	}
	fmt.Println(reply)
}
