package solution

/*
	leetcode: https://leetcode.com/problems/move-zeroes/
*/

func moveZeroes(nums []int) {
	k := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			k++
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
}
