package solution

import "sort"

/*
	leetcode: https://leetcode.com/problems/two-sum-less-than-k/
*/

func twoSumLessThanK(nums []int, k int) int {
	n := len(nums)
	if n < 2 {
		return -1
	}

	sort.Ints(nums)

	max := -1
	l, r := 0, n-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum < k {
			max = getMax(max, sum)
			l++
		} else {
			r--
		}

		l = getLeft(nums, l, r)
		r = getRight(nums, l, r)
	}

	return max
}

func getLeft(nums []int, l, r int) int {
	for l > 0 && l <= r {
		if nums[l-1] == nums[l] {
			l++
		} else {
			break
		}
	}
	return l
}

func getRight(nums []int, l, r int) int {
	n := len(nums)
	for r < n-1 && l <= r {
		if nums[r] == nums[r+1] {
			r--
		} else {
			break
		}
	}
	return r
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
