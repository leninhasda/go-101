package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

type master struct {
	user
}

func (m *master) notify() {
	fmt.Printf("sending master email to %s<%s>\n", m.name, m.email)
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

	u := user{
		name:  "user",
		email: "user@mail.com",
	}

	sendNotification(&u)

	// will work too
	// because the user.notify will be promoted
	sendNotification(&ad)

	m := master{
		user: user{
			name:  "master",
			email: "master@mail.com",
		},
	}

	sendNotification(&m)
}

func sendNotification(n notifier) {
	n.notify()
}
