package solution

import "math"

/*
	leetcode: https://leetcode.com/problems/minimum-path-sum/
*/

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := NewDP(m+1, n+1)
	dp[m][n-1], dp[m-1][n] = 0, 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = grid[i][j] + getMin(dp[i+1][j], dp[i][j+1])
		} // end loop
	} // end loop

	return dp[0][0]
}

func NewDP(m, n int) [][]int {
	a := make([][]int, m)
	for i := 0; i < len(a); i++ {
		a[i] = make([]int, n)
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = math.MaxInt32
		}
	}
	return a
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
