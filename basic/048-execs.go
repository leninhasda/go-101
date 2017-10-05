package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, lockErr := exec.LookPath("ls")
	if lockErr != nil {
		panic(lockErr)
	}
	// fmt.Println(binary, lookErr)
	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
