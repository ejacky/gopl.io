package ex4_4

import "fmt"

func ExampleRotate() {
	a := []int{0, 1, 2, 3, 4, 5}
	rotate(a, 2)
	fmt.Println(a)
	//output:
	//[2 3 4 5 0 1]

}
