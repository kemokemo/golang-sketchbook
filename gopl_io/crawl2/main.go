package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)
	var n int // the number of waiting to send worklist

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// this tokens limits the number of parallel crawler.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{} // get token
	list, err := Extract(url)
	<-tokens // relese token

	if err != nil {
		log.Printf("failed to extract url '%s': %v", url, err)
	}
	return list
}
