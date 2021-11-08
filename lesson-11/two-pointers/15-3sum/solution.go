package solution

import "sort"

/*
	leetcode: https://leetcode.com/problems/3sum
*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	sort.Ints(nums)

	var result [][]int
	for i := 0; i < len(nums); i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			}

			l = getLeft(nums, i, l, r)
			r = getRight(nums, i, l, r)
		} // end loop
	} // end loop

	return result
}

func getLeft(nums []int, i, l, r int) int {
	for l < r && l > i+1 && nums[l] == nums[l-1] {
		l++
	}
	return l
}

func getRight(nums []int, i, l, r int) int {
	for l < r && r < len(nums)-1 && nums[r] == nums[r+1] {
		r--
	}
	return r
}
