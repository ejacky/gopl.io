package ex4_4

func rotate(s []int, n int) {
	copy(s, append(s, s[:n]...)[n:])
}
