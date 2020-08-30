package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Extract returns extracted links from the root url of arg.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to GET from '%s', status code is %v, status is '%s'", url, resp.StatusCode, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse respose from '%s' to html document: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore the invalid URL
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
