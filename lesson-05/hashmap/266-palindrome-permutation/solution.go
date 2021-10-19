package solution

/*
	leetcode: https://leetcode.com/problems/palindrome-permutation/
*/

/*
	We will count frequency of each character in string.
	After that, we check how many character have frequency is odd.
	If len of string is even and don't have any odd : string can PermutePalindrome
	if len of string is odd and have only one odd: string can still PermutePalindrome
	Otherwise, string cannot PermutePalindrome

	Time complexity: O(n)
	Space complexity: O(1)
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
