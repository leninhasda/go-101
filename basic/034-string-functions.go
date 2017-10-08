package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("contains:	", s.Contains("test", "es"))
	p("Count:	", s.Count("test", "t"))
	p("HasPrefix:	", s.HasPrefix("test", "te"))
	p("HasSuffix:	", s.HasSuffix("test", "st"))
	p("Index:	", s.Index("test", "e"))
	p("Join:	", s.Join([]string{"e", "y", "e"}, "-"))
	p("Repeat:	", s.Repeat("test", 2))
	p("Replace:	", s.Replace("foo", "o", "0", -1))
	p("Replace:	", s.Replace("foo", "o", "0", 1))
	p("Split:	", s.Split("e-y-e", "-"))
	p("ToLower:	", s.ToLower("TesT#"))
	p("ToUpper:	", s.ToUpper("te0t"))

	p("len:	", len("hello	0"))
	p("char:	", "hello"[1])

	p("compare: ", s.Compare("Test", "test"))
	p("compare: ", s.Compare("test", "test"))
	p("equal fold: ", s.EqualFold("Test", "test"))
	p("equal fold: ", s.EqualFold("test", "testsdfasdf"))
	p("last index any: ", s.LastIndexAny("testes", "t"))
	p("last index: ", s.LastIndex("testes", "t"))
	p("title: ", s.Title("test"))
	p("title: ", s.Title("test text"))
	p("to title:", s.ToTitle("to title"))

}
