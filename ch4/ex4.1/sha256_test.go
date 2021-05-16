package ex4_1

import (
	"crypto/sha256"
	"fmt"
)

func ExampleDiff() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("%b\n%b\n", c1, c2)

	fmt.Println(Diff(&c1, &c2))

	// Output:
	// true
	// false
}
