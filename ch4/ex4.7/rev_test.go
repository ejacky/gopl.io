package ex4_7

import "fmt"

func ExampleReverse() {
	a := []rune{'a', '张', 'c', '~', 'e', 'f'}
	reverse(a)
	fmt.Printf("%c", a)
	//output:
	//[f e ~ c 张 a]

}
