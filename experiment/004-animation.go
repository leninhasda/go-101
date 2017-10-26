package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	sn := "=>"
	for i := 1; i <= 100; i++ {
		clear()
		for j := 0; j <= i; j++ {
			fmt.Print("-")
		}
		fmt.Print(sn)
		time.Sleep(time.Millisecond * 100)
	}
}
