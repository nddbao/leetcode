package solution

/*
	leetcode: https://leetcode.com/problems/sort-colors
*/

func sortColors(nums []int) {
	n := len(nums)
	k := 0

	// move 0 ahead
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}

	// move 1 ahead after 0
	for i := k; i < n; i++ {
		if nums[i] == 1 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}
}
