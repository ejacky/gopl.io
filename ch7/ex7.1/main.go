package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordsCounter int
type LinesCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {

	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)

	for input.Scan() {
		*c++
	}
	return int(*c), nil
}

func (c *LinesCounter) Write(p []byte) (int, error) {

	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanLines)

	for input.Scan() {
		*c++
	}
	return int(*c), nil
}

func main() {
	var w WordsCounter
	w.Write([]byte("nihao!\nhello world!"))
	fmt.Println(w)

	var l LinesCounter
	l.Write([]byte("nihao!\nhello world!"))
	fmt.Println(l)
}
