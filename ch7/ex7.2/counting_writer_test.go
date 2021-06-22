package ex7_2

import (
	"fmt"
)

func ExampleCountingWriter() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	CountingWriter(&c)
}
