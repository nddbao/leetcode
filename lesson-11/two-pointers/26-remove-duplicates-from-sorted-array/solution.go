package solution

/*
	leetcode: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/

/*
	 We will have two parts in our array:
		+ First: contains distinct number
		+ Second: contains elements that we have not visited yet.

	Iterating all of elements in array and using two pointers:
 		+ one pointer to keep track (lastIdx + 1) of our first part
		+ one pointer to keep track index of element we visit.
    Array is an non-decreasing, so if last item in first part is less than current visited element,
    We append current element by swapping value based on two pointers.

	Time Complexity: O(n)
	Space Complexity: O(1)

*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[k-1] {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}

	return k
}
