package solution

/*
	leetcode: https://leetcode.com/problems/subsets/
*/

func subsets(nums []int) [][]int {
	n := len(nums)
	numberOfSubsets := 1 << n

	result := make([][]int, numberOfSubsets)
	for i := 0; i < numberOfSubsets; i++ {
		sub := getSubset(nums, i)
		result[i] = sub
	}
	return result
}

func getSubset(nums []int, k int) []int {
	var a []int
	idx := 0
	for k != 0 {
		if k&1 == 1 {
			a = append(a, nums[idx])
		}
		k = k >> 1
		idx++
	}
	return a
}
