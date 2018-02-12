package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4

	v := reflect.ValueOf(x)
	fmt.Println("value:", v.Interface())
	// convert v => interface => assert to float64
	fmt.Println("type:", reflect.TypeOf(v.Interface().(float64)))
}
