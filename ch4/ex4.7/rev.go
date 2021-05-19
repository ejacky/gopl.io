package ex4_7

import "unicode/utf8"

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseUTF8(s []byte) {
	var size int
	for i := 0; i < len(s); {
		_, size = utf8.DecodeRune(s[i:])
		reverse(s[i : i+size])
		i += size
	}
	reverse(s)
}
