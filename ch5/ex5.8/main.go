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
	node := ElementByID(doc, "img_id")
	fmt.Println(node.Data)

	for _, a := range node.Attr {
		fmt.Println(a.Key, a.Val)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, startElement, id)
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre func(n *html.Node, id string) bool, id string) (ret *html.Node) {
	if pre != nil && pre(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = forEachNode(c, pre, id)
		if ret != nil {
			break
		}
	}

	return

	//if post != nil {
	//	post(n)
	//}
}

//!-forEachNode

//!+startend

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}

	return false
}

func endElement(n *html.Node) bool {

	return true
}

//!-startend
