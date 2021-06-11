// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	NetUrl "net/url"
	"os"
	"reflect"
	"strings"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(worklist []string) {
	var origin []string
	for _, oU := range worklist {
		origin = append(origin, GetDomain(oU))
	}

	fmt.Println(origin)

	//!+crawl
	var crawl = func(url string) []string {
		if exist, _ := InSlice(GetDomain(url), origin); exist {
			fmt.Println(url)
		}

		list, err := Extract(url, origin)
		if err != nil {
			log.Print(err)
		}
		return list
	}

	//!-crawl

	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, crawl(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(os.Args[1:])
}

//!-main

func Extract(url string, origin []string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
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
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
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

func InSlice(element interface{}, slice interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(element, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func GetDomain(url string) string {
	u, err := NetUrl.Parse(url)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(u.Hostname(), ".")
	if len(parts) < 2 {
		log.Fatal("invalid url :", url)
	}
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
	return domain
}
