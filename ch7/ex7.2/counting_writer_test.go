package ex7_2

import (
	"bytes"
	"fmt"
)

func ExampleCountingWriter() {
	wc, count := CountingWriter(&bytes.Buffer{})

	fmt.Println(*count)
	wc.Write([]byte("hello"))
	fmt.Println(*count)
	//output:
	//0
	//5
}
