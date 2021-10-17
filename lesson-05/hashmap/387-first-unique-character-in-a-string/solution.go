package solution

/*
	leetcode: https://leetcode.com/problems/first-unique-character-in-a-string/
*/

func firstUniqChar(s string) int {
	charArr := newCharArray()

	for i := 0; i < len(s); i++ {
		k := s[i] - 'a'
		if charArr[k] == -1 {
			charArr[k] = i
		} else {
			charArr[k] = -2
		}
	}

	min := len(s)
	for _, v := range charArr {
		if v < 0 {
			continue
		}

		if v < min {
			min = v
		}
	}

	if min == len(s) {
		return -1
	}

	return min
}

func newCharArray() [26]int {
	var a [26]int
	for i := 0; i < len(a); i++ {
		a[i] = -1
	}
	return a
}
