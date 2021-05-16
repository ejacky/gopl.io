package ex4_3

import "fmt"

func ExampleReverse() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)
	//output:
	//[5 4 3 2 1 0]

}
