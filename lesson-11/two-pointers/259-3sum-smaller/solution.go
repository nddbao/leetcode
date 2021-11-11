package solution

import "sort"

/*
	leetcode: https://leetcode.com/problems/3sum-smaller/
*/

/*
	We sort the array. Then we pick 1st and use 2-problem to solve.
	With 2-problem, we have left and right pointer to keep track.
		+ Left start from beginning and move to right side.
		+ Right start from the end and move to left side.

	There 2 two cases when calculating sum = 1st + nums[left] + nums[right]
		+ sum < target: we know all element in range [left+1 ... right] will be < target
			=> so we just add len of this range to our count and move left pointer to right side by one
		+ sum >= target: just move right pointer to left side by one.

	Time complexity: O(nlogn) + O(n^2) -> O(n^2)
	Space complexity: O(n) or O(1) (depend on sort)
*/
func threeSumSmaller(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	sort.Ints(nums)

	var count int
	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < target {
				count += r - l
				l++
			} else {
				r--
			}
		} // end loop
	} // end loop

	return count
}
