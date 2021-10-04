package solution

/*
	leetcode: https://leetcode.com/problems/happy-number/
*/

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	fast, slow := n, n
	for {
		fast = sumSquare(fast)
		fast = sumSquare(fast)
		slow = sumSquare(slow)

		if fast == 1 {
			return true
		}

		if fast == slow {
			return false
		}
	}
}

func sumSquare(k int) int {
	var s int
	for k != 0 {
		r := k % 10
		s += r * r
		k = k / 10
	}
	return s
}
