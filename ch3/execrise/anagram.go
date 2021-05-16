package anagram

//!+ 3.12
func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m := make(map[string]int)
	for i := range s1 {
		m[string(s1[i])] += 1
	}

	m1 := make(map[string]int)
	for i := range s2 {
		m1[string(s2[i])] += 1
	}

	for key, _ := range m {
		if m[key] != m1[key] {
			return false
		}
	}

	return true
}

//!- 3.12
