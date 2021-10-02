package solution

/*
	leetcode: https://leetcode.com/problems/power-of-two/
*/
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	if clearLastSetBit(n) == 0 {
		return true
	}

	return false
}

func clearLastSetBit(n int) int {
	return n & (n - 1)
}
