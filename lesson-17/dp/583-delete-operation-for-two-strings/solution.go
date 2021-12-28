package solution

/*
	leetcode: https://leetcode.com/problems/delete-operation-for-two-strings/
*/

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := NewDP(m+1, n+1)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = getMax(dp[i+1][j], dp[i][j+1])
			}
		} // end loop
	} // end loop

	return (m - dp[0][0]) + (n - dp[0][0])
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
