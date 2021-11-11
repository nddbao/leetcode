package solution

/*
	leetcode: https://leetcode.com/problems/subarray-product-less-than-k/
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	n := len(nums)

	var count int
	prod := 1
	left := 0
	for right := 0; right < n; right++ {
		prod *= nums[right]
		for prod >= k {
			prod /= nums[left]
			left++
		}
		count += right - left + 1
	}

	return count
}
