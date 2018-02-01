package main

import (
	"fmt"

	"github.com/leninhasda/gohttp"
)

func main() {
	req := gohttp.NewRequest()
	ch := make(chan *gohttp.Response)

	users := []string{"nahid", "shipu", "sujan"}

	for i := 0; i < len(users); i++ {
		req.
			FormData(map[string]string{"user": users[i]}).
			AsyncPost("https://httpbin.org/post", ch)
	}

	for i := 0; i < len(users); i++ {
		op := <-ch
		fmt.Println(op.GetBodyAsString())
		fmt.Println("====================")
	}
}
