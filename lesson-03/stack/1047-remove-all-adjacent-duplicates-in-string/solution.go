package solution

/*
	leetcode: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

*/

func removeDuplicates(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || (stack[len(stack)-1] != s[i]) {
			stack = append(stack, s[i])
			continue
		}

		stack = stack[:len(stack)-1]
	}
	return string(stack)
}
