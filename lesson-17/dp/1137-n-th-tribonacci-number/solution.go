package solution

/*
	leetcode: https://leetcode.com/problems/n-th-tribonacci-number/
*/

/*
	Time complexity: O(N)
	Space complexity: O(1)
*/
func tribonacci(n int) int {
	if n <= 1 {
		return n
	}

	if n == 2 {
		return 1
	}

	t0, t1, t2 := 0, 1, 1
	for i := 3; i <= n; i++ {
		tmp := t2
		t2 = t0 + t1 + t2
		t0, t1 = t1, tmp
	}

	return t2
}
