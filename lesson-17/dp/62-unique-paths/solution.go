package solution

/*
	leetcode: https://leetcode.com/problems/unique-paths/
*/

func uniquePaths(m int, n int) int {
	dp := NewDP(m+1, n+1)
	dp[m-1][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
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
