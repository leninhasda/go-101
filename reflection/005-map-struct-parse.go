package main

type Person struct {
	Name   string
	Age    int
	Emails []string
	Extra  map[string]string
}

func main() {

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
	err := Decode(input, &result)
	if err != nil {
		panic(err)
	}

	// spew.Dump(result)
}

func Decode(input interface{}, output interface{}) error {
	// inref, _ := input.(map[string]interface{})
	// outref := reflect.ValueOf(&output).Elem()
	// fmt.Println(reflect.ValueOf(reflect.ValueOf(&output).Elem().(Person)).String())
	// outref := reflect.TypeOf(&output).Elem()
	// typeOf := outref.Type()
	// fmt.Println("type:", outref.Kind())
	// for i := 0; i < outref.NumField(); i++ {
	// 	key := typeOf.Field(i).Name
	//
	// 	// switch outref.Field(i).Kind() {
	// 	// case reflect.Int:
	// 	outref.Field(i).Set(reflect.ValueOf(inref[key]))
	// 	// }
	// }
	return nil
}
