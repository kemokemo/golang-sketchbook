package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	exitOK = iota
	exitFail
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	var err, e error
	for _, arg := range args {
		e = createFile(arg)
		if e != nil {
			err = fmt.Errorf("%v:%v", err, e)
		}
	}

	if err != nil {
		fmt.Println("failed to create files: ", err)
		return exitFail
	}
	return exitOK

}

// createFile creates file. If the folder to create the file does not exist,
// create the folder beforehand and then create the file.
func createFile(arg string) error {
	dir := filepath.Dir(arg)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	f, err := os.Create(arg)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)

	}
	defer f.Close()

	return nil
}
