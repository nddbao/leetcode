package solution

/*
	leetcode: https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/

/*
	Using two pointers to maintain a sliding window that contain unique characters.
	We also use extra map to keep track character in window.
	Iterate each character in string, use this map to check current character already in window or not.
		+ If not, we delete s[left] character and move left ahead one step until current character not in map anymore.
		+ Otherwise, just add current character to map and update max value.

	Time complexity: O(n)
	Space complexity: O(n)
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
