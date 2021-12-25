package solution

/*
	leetcode: https://leetcode.com/problems/longest-increasing-subsequence/
*/

/*
	Time complexity: O(N^2)
	Space complexity: O(N)
*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	lis := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		lis[i] = 1
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				lis[i] = getMax(lis[i], lis[j]+1)
			}
		} // end loop
	} // end loop

	max := 0
	for i := 0; i < n; i++ {
		max = getMax(max, lis[i])
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
