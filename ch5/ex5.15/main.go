// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"errors"
	"fmt"
)

//!+
func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("params is empty")
	}

	var min int
	min = vals[0]
	for _, val := range vals {

		if min > val {
			min = val
		}
	}
	return min, nil
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("params is empty")
	}

	var max int
	max = vals[0]
	for _, val := range vals {

		if max < val {
			max = val
		}
	}
	return max, nil

}

//!-

func main() {
	//!+main
	fmt.Println(min())
	fmt.Println(min(3))
	fmt.Println(min(1, 2, 3, 4))
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(min(values...))
	//!-slice

	//!+main
	fmt.Println(max())
	fmt.Println(max(3))
	fmt.Println(max(-1, 2, 3, 4))
	//!-main

	fmt.Println("===========")

	//!+slice
	values = []int{1, 2, 3, 4}
	fmt.Println(max(values...))
	//!-slice
}
