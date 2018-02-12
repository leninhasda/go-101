package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x)) // float64
	// cause underling type is different
	fmt.Println("value:", reflect.ValueOf(x).String()) // value: <float64 value>

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind of v is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	type MyInt int
	var y MyInt = 7
	c := reflect.ValueOf(y)
	fmt.Println("type c:", c.Type())
	fmt.Println("value c:", c)
}
