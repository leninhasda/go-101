package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		fmt.Println("Timer 2 expired")
		<-timer2.C
	}()
	fmt.Println("end of line")
	time.Sleep(time.Second * 5) // wait until go routine stops

	// timer 2 expired wasn't pringing becase it was being stopped by following line
	// stop2 := timer2.Stop()
	// if stop2 {
	// 	fmt.Println("Timer 2 stopped")
	// }
}
