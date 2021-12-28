package solution

import "math"

/*
	leetcode: https://leetcode.com/problems/minimum-falling-path-sum/
*/

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := NewDP(n)
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			a := getDPValue(dp, i+1, j-1)
			b := getDPValue(dp, i+1, j)
			c := getDPValue(dp, i+1, j+1)
			dp[i][j] = matrix[i][j] + getMin(a, b, c)
		}
	}

	return getMin(dp[0]...)
}

func NewDP(n int) [][]int {
	a := make([][]int, n)
	for i := 0; i < len(a); i++ {
		a[i] = make([]int, n)
	}

	return a
}

func getDPValue(dp [][]int, i, j int) int {
	n := len(dp)
	if i == n {
		return 0
	}

	if i >= 0 && i < n && j >= 0 && j < n {
		return dp[i][j]
	}

	return math.MaxInt32
}

func getMin(a ...int) int {
	min := math.MaxInt32
	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}
