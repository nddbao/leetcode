package solution

func findMin(nums []int) int {
	l, h := initLowAndHigh(nums)
	if l > h {
		return nums[0]
	}

	if l == h || nums[l] < nums[h] {
		if nums[0] < nums[l] {
			return nums[0]
		}
		return nums[l]
	}

	for l <= h {
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

	return nums[0]
}

func initLowAndHigh(nums []int) (int, int) {
	l, h := 0, len(nums)-1
	for l < h {
		if nums[l] != nums[h] {
			break
		}
		h--
	}
	return l, h
}
