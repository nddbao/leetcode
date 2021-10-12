package solution

/*
	leetcode: https://leetcode.com/problems/kth-largest-element-in-an-array
*/

func findKthLargest(nums []int, k int) int {
	pivot := nums[len(nums)-1]
	p := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= pivot {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	nums[p], nums[len(nums)-1] = pivot, nums[p]

	if p > k-1 {
		return findKthLargest(nums[:p], k)
	}

	if p < k-1 {
		return findKthLargest(nums[p+1:], k-p-1)
	}

	return pivot
}
