package main

import "fmt"

type SSlice []string

// func (s *SSlice) pswap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

// this works becasue slice holds pointer to that slice/array
func (s SSlice) swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := SSlice{"z", "a", "x", "b", "c", "l", "h"}

	fmt.Println(s)
	s.swap(1, 2)
	fmt.Println(s)
	// s.pswap(1, 2)
	// fmt.Println(s)
}
