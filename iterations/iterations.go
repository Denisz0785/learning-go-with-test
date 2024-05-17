package iterations

func Repeat(a string, b int) (s string) {
	for i := 0; i < b; i++ {
		s += a
	}
	return s
}
