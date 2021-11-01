package solution

/*
	leecode: https://leetcode.com/problems/search-in-rotated-sorted-array/
*/

/*
	We can go through all of elements in array and search for target. But time complexity: O(n).

	Remembering that this is an rotated sorted array.
	So we can apply binary search with some modifications for updating low and high.
	Dividing array into two range: A= [low..mid] and B = [mid+1..high]

	We are sure that A or B have to be in order:
		+ If A is in order and target belongs to that range,
			update low and high to search target in A. Otherwise, we search B.
		+ Similarly, if B is in order and target belongs to that range
			update low and high to search target in B. Otherwise, we search A.

	Our final range will have only one element.
	Just check this element to know target in nums or not.

	Time complexity: O(log(n))
	Space complexity: O(1)

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
