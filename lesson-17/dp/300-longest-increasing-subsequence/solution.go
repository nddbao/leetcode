package solution

/*
	leetcode: https://leetcode.com/problems/longest-increasing-subsequence/
*/

/*
	We need lis (longest increasing subsequence) array to store max from index in to the end.
	We init it with 1 and go from the end to the beginning.
	For each iteration, we check all index after i and nums[i] < nums[j] and update lis[i]

	Finally, we just find max in lis.

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
