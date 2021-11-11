package solution

import "sort"

/*
	leetcode: https://leetcode.com/problems/4sum/
*/

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}

	sort.Ints(nums)

	var result [][]int
	for i1 := 0; i1 >= 0; i1 = getLeft(nums, i1+1, 0, n-4) {
		for i2 := i1 + 1; i2 >= 0; i2 = getLeft(nums, i2+1, i1+1, n-3) {
			sum1 := nums[i1] + nums[i2]
			l, r := i2+1, n-1
			for l < r {
				sum := sum1 + nums[l] + nums[r]
				if sum == target {
					result = append(result, []int{nums[i1], nums[i2], nums[l], nums[r]})
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}

				l = getLeft(nums, l, i2+1, r)
				if l < 0 {
					break
				}

				r = getRight(nums, r, l, n-1)
				if r < 0 {
					break
				}
			}

		} // end loop i2
	} // end loop i1

	return result
}

func getLeft(nums []int, i, min, max int) int {
	for i <= max {
		if i > min && nums[i-1] == nums[i] {
			i++
		} else {
			return i
		}
	}

	return -1
}

func getRight(nums []int, i, min, max int) int {
	for i >= min {
		if i < max && nums[i] == nums[i+1] {
			i--
		} else {
			return i
		}
	}

	return -1
}
