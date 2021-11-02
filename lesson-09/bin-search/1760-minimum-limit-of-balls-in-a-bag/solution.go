package solution

/*
	leetcode: https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/
*/

/*
	Note: remember to assign result := h
*/

func minimumSize(nums []int, maxOperations int) int {
	l, h := 1, int(1e9)
	result := h
	for l < h {
		mid := l + (h-l)/2
		if condition(nums, maxOperations, mid) {
			result = mid
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return result
}

func condition(nums []int, maxOperation, target int) bool {
	for _, v := range nums {
		q := (v - 1) / target
		maxOperation -= q

		if maxOperation < 0 {
			return false
		}
	}
	return true
}
