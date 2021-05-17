package ex4_6

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashes(strings []string) []string {

	s := ">+=@?"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	out := strings[:1]
	for i := 1; i <= len(strings)-1; i++ {
		if unicode.IsSpace([]rune(out[len(out)-1])) && unicode.IsSpace(strings[i]) {
			out = append(out, strings[i])
		}
	}
	return []string{}
}
