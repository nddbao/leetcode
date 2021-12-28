package solution

/*
	Time complexity: O(n*m)
	Space complexity: O(n*m)
*/

func longestCommonSubsequenceNonRecursive(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	dp := NewDP(m, n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = getMax(dp[i+1][j], dp[i][j+1])
			}
		} // end loop
	} // end loop

	return dp[0][0]
}

func NewDP(m, n int) [][]int {
	a := make([][]int, m+1)
	for i := 0; i < len(a); i++ {
		a[i] = make([]int, n+1)
	}
	return a
}
