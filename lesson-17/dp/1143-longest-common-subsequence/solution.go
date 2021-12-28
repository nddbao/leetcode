package solution

/*
	leetcode: https://leetcode.com/problems/longest-common-subsequence/
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	cache := NewCache(len(text1), len(text2))
	return helper(cache, text1, 0, text2, 0)
}

func helper(cache [][]int, text1 string, i int, text2 string, j int) int {
	if i == len(text1) || j == len(text2) {
		return 0
	}

	if cache[i][j] >= 0 {
		return cache[i][j]
	}

	var result int
	if text1[i] == text2[j] {
		result = 1 + helper(cache, text1, i+1, text2, j+1)
	} else {
		a := helper(cache, text1, i+1, text2, j)
		b := helper(cache, text1, i, text2, j+1)
		result = getMax(a, b)
	}

	cache[i][j] = result
	return result
}

func NewCache(n, m int) [][]int {
	a := make([][]int, n)
	for i := 0; i < len(a); i++ {
		a[i] = make([]int, m)
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = -1
		}
	}
	return a
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
