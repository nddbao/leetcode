package solution

/*
	leetcode: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[k-1] {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}

	return k
}
