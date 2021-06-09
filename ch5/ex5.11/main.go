// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	courses, circle := topoSort(prereqs)
	fmt.Println("has circle:", circle)

	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, bool) {
	var order []string
	var circle bool
	seen := make(map[string]bool)
	flag := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, course := range items {
			if has, ok := flag[course]; ok && !has {
				circle = true
			}

			if !seen[course] {
				seen[course] = true

				flag[course] = false
				visitAll(m[course])
				flag[course] = true

				order = append(order, course)
			}
		}
	}

	for course, _ := range m {
		visitAll([]string{course})
	}

	return order, circle
}

//!-main
