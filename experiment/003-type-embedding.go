package main

import "fmt"

type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		// this doesn't work
		// name:  "test 2",
		// email: "t@s.c",
		// must use like this
		user: user{
			name:  "test",
			email: "a@b.c",
		},
		level: "l",
	}
	fmt.Println(ad)

	// but this will work
	// as the embedded vars will be promoted
	fmt.Println(ad.name, ad.email)

}
