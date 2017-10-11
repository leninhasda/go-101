package main

import "fmt"

func main() {

	aa := []int{1, 2, 3, 4, 5, 6}
	bb := aa[3:6]
	// fmt.Println(cap(bb))
	bb[0] = 10
	bb = append(bb, 20)
	fmt.Println(cap(bb))
	bb = append(bb, 20)
	fmt.Println(cap(bb))
	bb = append(bb, 20)
	fmt.Println(cap(bb))
	bb = append(bb, 20)
	bb[0] = 99

	fmt.Println(aa, cap(aa))
	fmt.Println(bb, cap(bb))

}
