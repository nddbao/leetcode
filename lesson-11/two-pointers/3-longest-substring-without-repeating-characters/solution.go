package solution

/*
	leetcode: https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/

func lengthOfLongestSubstring(s string) int {
	distinct := make(map[byte]struct{})

	var l, max int
	for r := 0; r < len(s); r++ {
		for l < r {
			if _, ok := distinct[s[r]]; !ok {
				break
			}
			delete(distinct, s[l])
			l++
		}

		distinct[s[r]] = struct{}{}
		max = getMax(max, r-l+1)
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
