package solution

/*
	leetcode: https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/
*/

/*
	We find longest common subsequence that have maximum ascii value.
	After that, we compute ascii value of s1 and s2.
	Finally, we wil have minimum ascii delete sum:
		sumAscii(s1) - maxAsciLIS + sumAscii(s2) - maxAsciLIS

	Time complexity: O(m*n) where m is len of s1, n is len of s2
	Space complexity: O(m*n)
*/

func minimumDeleteSum(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)

	dp := NewDP(m+1, n+1)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = int(s1[i]) + dp[i+1][j+1]
			} else {
				dp[i][j] = getMax(dp[i+1][j], dp[i][j+1])
			}
		} // end loop
	} // end loop

	return sumASCII(s1) - dp[0][0] + sumASCII(s2) - dp[0][0]
}

func NewDP(m, n int) [][]int {
	a := make([][]int, m)
	for i := 0; i < len(a); i++ {
		a[i] = make([]int, n)
	}
	return a
}

func sumASCII(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i])
	}
	return sum
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
