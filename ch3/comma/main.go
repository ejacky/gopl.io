// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma3(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-

//!+ 3.10
func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	pre := n % 3

	var buf bytes.Buffer
	buf.Write([]byte(s[:pre]))
	for i := pre; i < n; i += 3 {
		buf.Write([]byte(","))
		buf.Write([]byte(s[i : i+3]))
	}
	return buf.String()
}

//!-

//!+ 3.11
func comma3(s string) string {
	var buf bytes.Buffer

	var start, end, frac int

	if s[0] == '+' || s[0] == '-' {
		buf.Write([]byte(string(s[0])))
		start = 1
	}

	n := len(s)
	index := strings.Index(s, ".")

	if index == -1 {
		end = n
		frac = n - 1
	} else {
		end = index - 1
		frac = index + 1
	}
	floatStr := comma2(s[start:end])

	buf.Write([]byte(floatStr))

	if frac != n-1 {
		buf.Write([]byte("."))
		buf.Write([]byte(s[frac:]))
	}

	return buf.String()
}

//!-
