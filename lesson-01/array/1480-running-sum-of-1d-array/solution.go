package solution

/*
	leetcode: https://leetcode.com/problems/running-sum-of-1d-array/
*/

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}
