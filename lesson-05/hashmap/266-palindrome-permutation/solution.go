package solution

/*
	leetcode: https://leetcode.com/problems/palindrome-permutation/
*/

func canPermutePalindrome(s string) bool {
	var charArray [26]int

	for i := 0; i < len(s); i++ {
		charArray[s[i]-'a']++
	}

	odd := 0
	for _, v := range charArray {
		if v%2 != 0 {
			odd++
		}
	}

	if len(s)%2 == 0 && odd == 0 {
		return true
	}

	if len(s)%2 != 0 && odd == 1 {
		return true
	}

	return false
}
