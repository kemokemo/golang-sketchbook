package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	fmt.Println("Hello, WebAssembly!")
	return 0
}
