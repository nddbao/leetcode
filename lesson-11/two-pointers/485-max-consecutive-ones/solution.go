package solution

/*
	leetcode: https://leetcode.com/problems/max-consecutive-ones/
*/

func findMaxConsecutiveOnes(nums []int) int {
	var max int

	left, right := -1, 0
	for ; right < len(nums); right++ {
		if nums[right] == 1 {
			if left < 0 {
				left = right
			}
			max = getMax(max, right-left+1)
			continue
		}
		left = -1
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
