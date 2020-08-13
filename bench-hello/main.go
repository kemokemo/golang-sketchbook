package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	wg := &sync.WaitGroup{}
	var e, err error

	wg.Add(1)
	go func() {
		e = heavyFunc()
		if e != nil {
			err = fmt.Errorf("%v, %v", err, e)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		e = simpleFunc()
		if e != nil {
			err = fmt.Errorf("%v, %v", err, e)
		}
		wg.Done()
	}()

	wg.Wait()

	if err != nil {
		fmt.Println("failed to work: ", err)
		return exitCodeFailed
	}
	return exitCodeOK
}

func heavyFunc() error {
	time.Sleep(time.Millisecond * 50)
	return nil
}

func simpleFunc() error {
	time.Sleep(time.Microsecond * 1)
	return nil
}
