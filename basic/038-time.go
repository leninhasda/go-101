package main

import (
	"fmt"
	"strings"
	"time"
)

func format(f string, d time.Time) string {
	mp := map[string]string{
		"d": "02",
		"D": "Mon",
		"j": "2",
		"l": "Monday",

		"F": "January",
		"m": "01",
		"M": "Jan",
		"n": "1",

		"Y": "2006",
		"y": "06",

		"A": "PM",
		"a": "pm",
		"g": "3",
		"G": "15",
		"h": "3",
		"H": "15",
		"i": "04",
		"s": "05",
	}

	for k, v := range mp {
		f = strings.Replace(f, k, v, 1)
	}
	println(f)
	return d.Format(f)
}

func main() {
	p := fmt.Println

	now := time.Now()

	p(now.Format("Mon Jan 2 15:04:05 MST 2006"))
	p(now.Format("06-01-02 January PM"))

	// p(now.Format(format("Y-m-d")))
	p(format("Y-m-d H:i:s", now))
}
