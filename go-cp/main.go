package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	exitOK = iota
	exitFail
)

func main() {
	os.Exit(run(os.Args))
}

// run is an endpoint of this app
func run(args []string) int {
	if len(args[1:]) < 2 {
		fmt.Println("please specify 'src' and 'destination'")
		return exitFail
	}

	err := copyFile(args[1], args[2])
	if err != nil {
		fmt.Println("failed to copy file: ", err)
		return exitFail
	}

	return exitOK
}

func copyFile(src, dst string) (err error) {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read source file: %v", err)
	}

	err = ioutil.WriteFile(dst, b, 0640)
	if err != nil {
		return fmt.Errorf("failed to write destination file: %v", err)
	}

	return nil
}
