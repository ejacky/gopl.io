package main

import "fmt"

func main() {
	fmt.Println(fakeRet())
}

func fakeRet() (err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil: // no panic
		case bailout{}: // "expected" panic
			err = fmt.Errorf("mock no return")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()

	panic(bailout{})
}
