package main

import (
	"fmt"
	"sort"
)

type ReverseSrting []string

func (s ReverseSrting) Len() int {
	return len(s)
}
func (s ReverseSrting) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ReverseSrting) Less(i, j int) bool {
	return s[i] > s[j]
}

func main() {
	strs := []string{"b", "z", "a"}

	// var strcopy []string // not work

	// strcopy := make([]string, len(strs)) // works
	// copy(strcopy, strs)

	strcopy := make([]string, 0) // work
	strcopy = append(strcopy, strs...)

	// strcopy := copy(strs, strs)strs

	sort.Strings(strs) // sort strs in assending

	istrs := sort.StringSlice(strs) // sort and return slice

	// sort.Sort(sort.Reverse(sort.StringSlice(strcopy))) // reverse sort
	sort.Sort(ReverseSrting(strcopy)) // custom reverse sorter
	// fmt.Printf("%T\n", istrs)

	// var x []string
	// x = istrs.([]string)
	// fmt.Printf("%T", x)
	fmt.Println(strs)
	fmt.Println(istrs)
	fmt.Println(strcopy)
}
