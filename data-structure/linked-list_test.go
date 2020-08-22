package ds_test

import (
	"testing"

	ds "github.com/leninhasda/go-101/data-structure"
)

func Test_LLAppend(t *testing.T) {
	list := new(ds.Node)
	list.Append(10)
	list.Append(20)
	if list.Len() != 2 {
		t.Error("error")
	}
	// spew.Dump(list)
}

func Test_LLFind(t *testing.T) {
	list := new(ds.Node)
	list.Append(10)
	list.Append(20)
	list.Append(30)
	if list.Find(20) != true {
		t.Error("error")
	}
}

func Test_LLRemove(t *testing.T) {
	list := new(ds.Node)
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	if list.Len() != 4 {
		t.Error("error")
	}
	list.Remove(20)
	if list.Len() != 3 {
		t.Error("error")
	}
}
