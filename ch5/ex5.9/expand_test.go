package ex5_9

import (
	"testing"
)

func TestExpend(t *testing.T) {

	if expand("", retBar) != "" {
		t.Error(`expand("", retBar) !=  ""`)
	}

	if expand("abc", retBar) != "abc" {
		t.Error(`expand("abc", retBar) !=  "abc"`)
	}

	if expand("c$ad", retBar) != "cbar" {
		t.Error(`expand("c$ad", retBar) !=  "cbar"`)
	}

	if expand("$a$b", retBar) != "barbar" {
		t.Error(`expand("$a$b", retBar) !=  "barbar"`)
	}
}

func retBar(s string) string {
	return "bar"
}
