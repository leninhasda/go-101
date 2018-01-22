package main

import (
	"log"
	"sync"
	"time"

	"github.com/leninhasda/go-101/concurrency-pattern/work"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{name}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}

	}

	wg.Wait()
	p.Shutdown()
}
