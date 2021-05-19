package ex4_7

import "fmt"

func ExampleReverse() {
	a := []byte("a张c~ef")
	reverseUTF8(a)
	fmt.Printf("%s", a)
	//output:
	//fe~c张a

}
