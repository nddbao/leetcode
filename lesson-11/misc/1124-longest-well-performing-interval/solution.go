package solution

/*
	leetcode: https://leetcode.com/problems/longest-well-performing-interval/
*/

/*
	We will check longest WPI from n to 1 util we get one.
	For checking WPI with k len:
		+ we use sliding window with k len
		+ Move from start of array util we cannot move
		+ If on the way, we found a window satisfied -> just return true.

	Time complexity: O(n^2)
	Space complexity: O(1)
	TODO: improve time complexity
*/
func longestWPI(hours []int) int {
	n := len(hours)
	for i := n; i >= 1; i-- {
		if checkWPI(hours, i) {
			return i
		}
	}

	return 0
}

func checkWPI(hours []int, k int) bool {
	var count int
	for i := 0; i < k; i++ {
		if isTiringDay(hours[i]) {
			count++
		}
	}

	if isWPI(k, count) {
		return true
	}

	for i := k; i < len(hours); i++ {
		if isTiringDay(hours[i-k]) {
			count--
		}
		if isTiringDay(hours[i]) {
			count++
		}

		if isWPI(k, count) {
			return true
		}
	}
	return false
}

func isWPI(n, tiringDays int) bool {
	return n-tiringDays < tiringDays
}

func isTiringDay(val int) bool {
	return val > 8
}
