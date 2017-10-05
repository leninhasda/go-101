package main

import (
	"os/exec"
	"strings"
	"time"
)

func clear() {
	// cls, _ := exec.LookPath("clear")
	//
	// argv := []string{}
	// env := os.Environ()
	//
	// er := syscall.Exec(cls, argv, env)
	// if er != nil {
	// 	panic(er)
	// }
	cls := exec.Command("clear")
	cls.Output()

}

sigs := make(chan os.Signal, 1)
done := make(chan bool, 1)

signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

go func() {
	sig := <-sigs
	fmt.Println()
	fmt.Println(sig)
	done <- true
}()

fmt.Println("awaiting")
<-done
fmt.Println("exiting")

func main() {
	for i := 0; i < 100; i++ {
		clear()
		sp := strings.Repeat(" ", i)
		print(sp)
		print(">")
		time.Sleep(time.Second)
	}
}
