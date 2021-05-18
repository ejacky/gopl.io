package ex4_6

import (
	"unicode"
)

func squashes(string []byte) []byte {

	out := string[:0]
	for i := 0; i <= len(string)-1; i++ {
		if len(out) == 0 && unicode.IsSpace(rune(string[i])) {
			out = append(out, ' ')
			continue
		}
		if unicode.IsSpace(rune(string[i])) && unicode.IsSpace(rune(out[len(out)-1])) {
			continue
		}

		if unicode.IsSpace(rune(string[i])) {
			out = append(out, ' ')
		} else {
			out = append(out, string[i])
		}
	}
	return out
}
