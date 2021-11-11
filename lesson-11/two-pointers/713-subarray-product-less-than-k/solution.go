package solution

/*
	leetcode: https://leetcode.com/problems/subarray-product-less-than-k/
*/

/*
	We use two pointers [left .. right] to keep track range.
	That range will has  product but less than K.
	Left and right pointer will move to right side

	First, we set left and right at the begin.
	Then we  have a loop to move right by one:
		+ Calculate new product.
		+ If prod >=k, we have to move left by one until we have prod < k.
		+ Then, we will add len(range) to our result.

	Time complexity: O(n)
	Space complexity: O(1)
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
