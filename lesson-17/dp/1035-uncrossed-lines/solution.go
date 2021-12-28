package solution

/*
	leetcode: https://leetcode.com/problems/uncrossed-lines/
*/

/*
	Similar to longest common subsequence.

	Time complexity: O(n*m) where m is len of nums1, n is len of nums2
	Space complexity: O(n*m)
*/

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)

	dp := NewDP(m+1, n+1)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = getMax(dp[i+1][j], dp[i][j+1])
			}
		} // end loop
	} // end loop

	return dp[0][0]
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
