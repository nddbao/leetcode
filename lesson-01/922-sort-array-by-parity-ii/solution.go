package solution

/*
	leetcode: https://leetcode.com/problems/sort-array-by-parity-ii/
*/

func sortArrayByParityII(nums []int) []int {
	k := 0
	for i := 0; i < len(nums); i += 2 {
		if nums[i]%2 == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k += 2
		}
	}

	for i := 1; i < len(nums); i += 2 {
		if nums[i]%2 == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k += 2
		}
	}

	return nums
}
