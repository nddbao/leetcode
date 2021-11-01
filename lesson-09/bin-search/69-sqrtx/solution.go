package solution

/*
	leetcode: https://leetcode.com/problems/sqrtx/
*/

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	var result int
	l, h := 1, x/2
	for l <= h {
		mid := l + (h-l)/2
		exp := mid * mid

		if exp == x {
			return mid
		}

		if exp > x {
			h = mid - 1
			continue
		}

		result = mid
		l = mid + 1
	}
	return result
}
