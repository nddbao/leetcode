package solution

import "sort"

/*
	leetcode: https://leetcode.com/problems/two-sum-less-than-k/
*/

/*
	First we sort array and apply two pointers (left, right)
     + Left start from beginning:  move to right side.
	 + Right start from the end: move to left side.

	We have max to keep track our largest sum < k.
	For each time, we check sum = nums[left] + nums[right]
	We divide it into two cases:
		+ sum < k: just update max and move left
		+ sum >= k: move right


	Time complexity: O(nlogn) + O(n) -> O(nlogn)
	Space complexity: O(n) or O(1) depend on sort
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
