package solution

func findMin(nums []int) int {
	if len(nums) == 1 || nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}

	l, h := 0, len(nums)-1
	for l < h {
		mid := l + (h-l)/2

		if nums[mid] <= nums[h] {
			h = mid
			continue
		}

		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		l = mid + 1
	}

	return 0
}
