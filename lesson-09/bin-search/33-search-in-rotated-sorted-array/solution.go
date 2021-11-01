package solution

/*
	leecode: https://leetcode.com/problems/search-in-rotated-sorted-array/
*/

func search(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l < h {
		mid := l + (h-l)/2

		// check range: [l...mid]
		if isAscendingOrder(nums[l : mid+1]) {
			if nums[l] <= target && target <= nums[mid] {
				h = mid
			} else {
				l = mid + 1
			}
			continue
		}

		// check range: [mid+1...h]
		if nums[mid+1] <= target && target <= nums[h] {
			l = mid + 1
		} else {
			h = mid
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}

func isAscendingOrder(nums []int) bool {
	return nums[0] <= nums[len(nums)-1]
}
