package solution

/*
	leetcode: https://leetcode.com/problems/valid-parentheses/
*/

/*
	Using stack DS to solve this problem. Our stack will keep only open brackets.
	When iterating every characters in string, we have two cases:
		+ First case: if opening, just put it into stack
		+ Second case:  if closing, get opening from top  stack.
						If we don't have any or they is not the same type.
						We know for sure string is not a valid parentheses.

	When done for loop, just checking stack is empty or not.
	If empty, we know all parentheses matching each other in correct order.

	Time Complexity: O(n)
	Space Complexity: O(n)
*/

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

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
