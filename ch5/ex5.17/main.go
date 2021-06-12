// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	file, _ := os.Open("t.html")
	defer file.Close()

	doc, _ := html.Parse(file)

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	for _, image := range images {
		fmt.Println(image.Data)
	}

	fmt.Println("===========")

	for _, heading := range headings {
		fmt.Println(heading.Data)
	}
}

func ElementsByTagName(doc *html.Node, tags ...string) []*html.Node {
	var nodes []*html.Node
	if len(tags) == 0 {
		return nodes
	}

	for _, tag := range tags {
		nodes = forEachNode(doc, tag, nodes)
	}

	return nodes

}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, tag string, nodes []*html.Node) []*html.Node {

	if n.Type == html.ElementNode && n.Data == tag {
		nodes = append(nodes, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = forEachNode(c, tag, nodes)
	}

	return nodes

}

//!-forEachNode
