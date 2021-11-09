package solution

/*
	leetcode: https://leetcode.com/problems/maximum-score-from-removing-stones/
*/

func maximumScore(a int, b int, c int) int {
	a, b, c = sortValues(a, b, c)
	max := 0
	for b > 0 {
		a, b, max = a-1, b-1, max+1
		a, b, c = sortValues(a, b, c)
	}

	return max
}

func sortValues(a, b, c int) (int, int, int) {
	if a < b {
		a, b = b, a
	}

	if c >= a {
		return c, a, b
	} else if c >= b {
		return a, c, b
	}
	return a, b, c
}
