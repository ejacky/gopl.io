package anagram

import "testing"

func TestIsAnagram(t *testing.T) {

	if !isAnagram("hello", "olleh") {
		t.Error(`isAnagram("hello", "olleh") = false`)
	}

	if isAnagram("hello", "elleh") {
		t.Error(`isAnagram("hello", "olleh") = true`)
	}

	if !isAnagram("A man, a plan, a canal: Pana2a", "A 2an, a plan, a canal: Panama") {
		t.Error(`isAnagram("A man, a plan, a canal: Pana2a", "A 2an, a plan, a canal: Panama") = false`)
	}
}
