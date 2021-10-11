package solution

/*
	leetcode: https://leetcode.com/problems/valid-parentheses/
*/

func isValid(s string) bool {
	m := map[byte]byte{']': '[', '}': '{', ')': '('}

	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		v, ok := m[c]
		if !ok {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 || v != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
