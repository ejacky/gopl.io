package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) { /* ... */
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode && !(n.Parent.Type == html.ElementNode && (n.Parent.Data == "style" || n.Parent.Data == "script") && len(n.Data) != 0) {
		str := strings.TrimSpace(strings.Trim(n.Data, "\n"))
		if len(str) > 0 {

			input := bufio.NewScanner(strings.NewReader(str))
			input.Split(bufio.ScanWords)

			for input.Scan() {
				words++
			}

			fmt.Println(str)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		newWords, newImages := countWordsAndImages(c)
		words += newWords
		images += newImages
	}

	return
}

func main() {
	for _, url := range os.Args[1:] {
		words, images, _ := CountWordsAndImages(url)
		fmt.Printf("words:%d\timages:%d\n", words, images)
	}
}
