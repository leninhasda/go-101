package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// v.SetFloat(7.1) // Error: will panic.
	// cause x is a copy of original value
	fmt.Println("can set address v:", v.CanSet())

	vv := reflect.ValueOf(&x).Elem()
	// vv type is *float64
	fmt.Println("can set address vv:", vv.CanSet())
}
