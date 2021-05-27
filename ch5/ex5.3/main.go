package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

var specialElem = 0

func main() {
	counts := make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	printText(counts, doc)

}

func printText(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script") {
		specialElem = 1
		fmt.Print("tag:")
		fmt.Println(n.Data)
	}
	if n.Type == html.TextNode && specialElem != 1 {
		fmt.Println(strings.TrimSpace(strings.Trim(n.Data, "\n")))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printText(counts, c)
	}
}
