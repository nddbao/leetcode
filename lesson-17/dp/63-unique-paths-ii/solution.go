package solution

/*
	leetcode: https://leetcode.com/problems/unique-paths-ii/
*/

/*
	Similar to unique-paths, but at index (i, j) that is obstacle we will set value equal to 0.

	Time complexity: O(m*n)
	Space complexity: O(m*n)
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := NewDP(m+1, n+1)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = 1
				continue
			}

			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i][j+1] + dp[i+1][j]
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
