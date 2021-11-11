package solution

import "sort"

/*
	leetcode: https://leetcode.com/problems/3sum-smaller/
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
