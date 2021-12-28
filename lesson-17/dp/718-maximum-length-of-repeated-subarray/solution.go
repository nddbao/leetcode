package solution

/*
	leetcode: https://leetcode.com/problems/maximum-length-of-repeated-subarray/
*/

func findLength(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := NewDP(m+1, n+1)
	max := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				max = getMax(max, dp[i][j])
			}
		} // end loop
	} // end loop

	return max
}

func NewDP(m, n int) [][]int {
	a := make([][]int, m)
	for i := 0; i < len(a); i++ {
		a[i] = make([]int, n)
	}

	return a
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
