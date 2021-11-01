package solution

/*
	leetcode: https://leetcode.com/problems/sqrtx/
*/

/*
	We can check number from 1 to x to find out result. But running time will be: O(n).
	By using binary search, we can reduce running time to O(log(n))
	Search from range [low .. high] with init value (low = 1, high = x/2)
	We have mid = (low + high)/2 and calculate square of mid: exp = mid * mid
	 + if exp == x: we just return mid
	 + if exp > x: update high = mid -1
     + otherwise: update result = mid and low = mid + 1
	Keep continue when low still less or equal to high.

	Time complexity: O(log(n))
	Space complexity: O(1)
*/

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	var result int
	l, h := 1, x/2
	for l <= h {
		mid := l + (h-l)/2
		exp := mid * mid

		if exp == x {
			return mid
		}

		if exp > x {
			h = mid - 1
			continue
		}

		result = mid
		l = mid + 1
	}
	return result
}
