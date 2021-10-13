package solution

/*
	leetcode: https://leetcode.com/problems/kth-largest-element-in-an-array
*/

/*
	We can use QuickSelect algorithm to slow this.
	Dived array into 2 partitions: one is greater than pivot, one is less than or equal to pivot
	We have look like this:
	[0...p-1][p][p+1...len(a)-1]
	If p == k: we know for sure is the k-th largest element (k-th 0-based) nums[idx]
	If p > k: we will recurse to left partition
	If p < k: we recurse to right partition

	Time: Worst Case O(n^2), Average O(n), Best Case: O(n)

	Note: ask k-th is 0-based or 1-based???
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
