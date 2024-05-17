package slice

func Sum(arr [5]int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

func SumSlice(src []int) (sum int) {
	for _, v := range src {
		sum += v
	}
	return
}

func SumAll(src1 ...[]int) []int {

	var sums []int
	for _, v := range src1 {
		sums = append(sums, SumSlice(v))
	}
	return sums
}

func SumTails(src ...[]int) (sums []int) {
	for _, v := range src {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, SumAll(tail)...)
		}
	}
	return sums
}
