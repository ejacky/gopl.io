package ex4_1

func Diff(sha1, sha2 *[32]byte) int {

	c := 0
	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			if sha1[i]>>j&1 != sha2[i]>>j&1 {
				c++
			}
		}
	}
	return c
}
