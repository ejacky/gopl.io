package ex4_5

func eliminate(strings []string) []string {

	out := strings[:1]
	for i := 1; i <= len(strings)-1; i++ {
		if out[len(out)-1] != strings[i] {
			out = append(out, strings[i])
		}
	}
	return out

}
