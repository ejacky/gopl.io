package ex4_5

import "fmt"

func ExampleEliminate() {
	data := []string{"one", "one", "three", "three", "four", "five"}
	fmt.Println(eliminate(data))
	//output:
	//[one three four five]

}
