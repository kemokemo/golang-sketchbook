package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Hello, Goroutine")
		ch <- "Hey!"
	}()

	// stop until receiving the string
	str := <-ch
	fmt.Println(str)
}
