package solution

/*
	leetcode: https://leetcode.com/problems/first-unique-character-in-a-string/
*/

/*
	We can use an array charArr that represents for 26 alphabet. Init value will be -1.
	This array will store the index that character appears in string.
	For example:
		charArray [-1 -1 .. -1]
		idx = character - 'a' -> value of 'c' will have idx 2

	We go through string and check current character is set or not in charArray.
	If not, we will set the its index. Otherwise, we set to -2.
	For example:
	 	S = "acabdd"
					  'a' 'b' 'c' 'd' ....
		=> charArray [ -2  3   1   -2	-1  ... -1]

	We know that character appears 1 time in string if and only if value in charArray >= 0
	Just checking charArray with that constraint and find MIN.
	MIN is our result.

	Time complexity: O(n)
	Space complexity: O(1)


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
