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
	if n.Type == html.TextNode && !(n.Parent.Type == html.ElementNode && (n.Parent.Data == "style" || n.Parent.Data == "script") && len(n.Data) != 0) {
		str := strings.TrimSpace(strings.Trim(n.Data, "\n"))
		if len(str) > 0 {
			fmt.Println(str)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printText(counts, c)
	}
}
