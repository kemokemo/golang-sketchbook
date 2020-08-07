package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fd := int(os.Stdin.Fd())
	if terminal.IsTerminal(fd) {
		fmt.Println("not pipe, fd is ", fd)
	} else {
		b, _ := ioutil.ReadAll(os.Stdin)
		fmt.Printf("pipe, fd is %v, %v", fd, string(b))
	}
}
