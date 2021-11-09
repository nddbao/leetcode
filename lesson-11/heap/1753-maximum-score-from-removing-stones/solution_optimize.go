package solution

func maximumScoreOptimize(a int, b int, c int) int {
	a, b, c = sortValues(a, b, c)
	if a > b+c {
		return b + c
	}

	return a + (b+c-a)/2
}
