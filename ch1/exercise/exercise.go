package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//!+exercise 1.1
	fmt.Println(os.Args[0])
	//!-exercise 1.1

	//!+exercise 1.2
	for index, value := range os.Args[1:] {
		fmt.Println(index, value)
	}
	//!-exercise 1.2

}

//!+exercise 1.3
func echo1(elems []string, sep string) string {
	var s string
	for i := 1; i < len(elems); i++ {
		s += sep + elems[i]
		sep = " "
	}
	return s
}

func echo2(elems []string, sep string) string {
	s, sep := "", ""
	for _, arg := range elems[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3(elems []string, sep string) string {
	return strings.Join(os.Args[1:], " ")
}

//!-exercise 1.3
