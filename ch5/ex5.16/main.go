package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(StringJoin(",", "aa", "bb"))
}

func StringJoin(sep string, elems ...string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	var buf bytes.Buffer
	buf.WriteString(elems[0])
	for _, s := range elems[1:] {
		buf.WriteString(sep)
		buf.WriteString(s)
	}
	return buf.String()
}
