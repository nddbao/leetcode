package solution

/*
	leetcode: https://leetcode.com/problems/sort-array-by-parity/
*/

func sortArrayByParity(nums []int) []int {
	k := -1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			k++
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
	return nums
}
