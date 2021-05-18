package ex4_6

import "fmt"

func ExampleSquashes() {
	data := []byte{'z', ' ', 's', '\t', '\n', 'n', 'w'}
	fmt.Printf("%s", squashes(data))
	//output:
	//z s nw

}
