package ex4_7

func reverse(s []rune) {
	//for item := range s {
	//	fmt.Println(item)
	//}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
