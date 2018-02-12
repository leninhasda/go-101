package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
)

func main() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	result := Person{}
	result.Name = "hello"
	result.Age = 10
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	spew.Dump(result)
}
