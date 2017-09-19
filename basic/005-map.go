package main

import "fmt"

type Map map[string]interface{}

func main() {

	mp := Map{
		"key1": 10,
		"key2": "string type",
		"key3": true,
	}

	fmt.Println(mp)

	mp2 := map[string]interface{}{
		"foo":  1,
		"bar":  "it is",
		"test": 10.1,
	}
	delete(mp2, "foo")
	fmt.Println(mp2)
}
