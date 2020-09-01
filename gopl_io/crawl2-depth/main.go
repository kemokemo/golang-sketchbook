package main

import (
	"flag"
	"fmt"
	"log"
)

var depth = flag.Int("depth", 1, "depth to search links")

// Links is the links found on the pages with deapth.
type Links struct {
	depth int
	links []string
}

func init() {
	flag.Parse()
}

func main() {
	worklist := make(chan Links)
	var n int // the number of waiting to send worklist

	n++
	go func() { worklist <- Links{depth: 0, links: flag.Args()} }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth > *depth {
			continue
		}
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					links := crawl(link)
					worklist <- Links{depth: list.depth + 1, links: links}
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
