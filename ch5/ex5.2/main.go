package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

//!+
func main() {
	counts := make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	countElem(counts, doc)

	fmt.Printf("element\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
}

func countElem(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countElem(counts, c)
	}
}
