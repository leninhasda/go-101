package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func ShellExec(arg string) (string, string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	args := strings.Split(arg, " ")
	cmd := exec.Command(args[0], args[1:]...)

	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		panic(err)
	}
	err_str, _ := ioutil.ReadAll(stderr)
	out_str, _ := ioutil.ReadAll(stdout)

	return string(out_str), string(err_str)
}

func main() {
	out, err := ShellExec("python p.py")

	fmt.Printf("out %s\n\n", out)
	fmt.Printf("err %s\n\n", err)
}
