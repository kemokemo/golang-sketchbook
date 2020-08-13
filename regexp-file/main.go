package main

import (
	"fmt"
	"os"
	"regexp"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	pattern := `test.*.txt`
	reg := regexp.MustCompile(pattern)
	fmt.Printf("regexp string is '%s'\n", pattern)

	fName := "test.txt"
	result := reg.MatchString(fName)
	fmt.Printf("file name: %s, matched: %v\n", fName, result)

	fName = "test_01.txt"
	result = reg.MatchString(fName)
	fmt.Printf("file name: %s, matched: %v\n", fName, result)

	return exitCodeOK
}
