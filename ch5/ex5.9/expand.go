package ex5_9

import "strings"

func expand(s string, f func(string) string) string {
	if strings.Contains(s, "$") {
		count := strings.Count(s, "$")
		var newS string
		var index int

		index = strings.Index(s, "$")
		newS += s[0:index]

		tmp := make([]string, count)
		tmp[0] += s[index+1:]
		for i := 0; i < count; i++ {
			index = strings.Index(tmp[i], "$")
			if index == -1 {
				newS += f(tmp[i])
			} else {
				newS += f(tmp[i][0:index])
				tmp[i] += s[index+1:]
			}
		}

		return newS
	}
	return s
}
