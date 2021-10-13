package solution

import (
	"strconv"
)

/*
	leetcode: https://leetcode.com/problems/evaluate-reverse-polish-notation/
*/

/*
	RPN format : number number operator -> result = number operator number
	We use stack DS. Iterate the array and push number to stack.
	If we meet a operator,  we will get two number from stack and calculate result. T
	Then we add result to stack
	Finally we have one left in our stack. This is our result.

	Time Complexity: O(n)
	Space Complexity: O(n)

	Note: check the order of numbers when doing operator
*/

func evalRPN(tokens []string) int {
	operatorFns := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	var stack []int
	for _, s := range tokens {
		fn, ok := operatorFns[s]
		if !ok {
			v, _ := strconv.Atoi(s)
			stack = append(stack, v)
			continue
		}

		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack[len(stack)-2] = fn(a, b)
		stack = stack[:len(stack)-1]
	}

	return stack[0]
}
