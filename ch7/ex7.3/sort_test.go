// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package ex7_3

import (
	"fmt"
	treesort "gopl/ex7.3"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func ExampleTree_String() {

	tree := new(tree)

	add(tree, 2)
	add(tree, 3)
	fmt.Println(tree.String())
	add(tree, 4)
	fmt.Println(tree)
	//output:
	//[0, 2, 3]
	//[0, 2, 3, 4]

}
