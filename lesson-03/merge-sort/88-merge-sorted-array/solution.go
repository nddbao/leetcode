package solution

/*
	leetcode: https://leetcode.com/problems/merge-sorted-array/
*/

/*
	1. At first we cannot create new array, and compare element between two of array.
	Then we can override nums1. Time: O(n+m), space: O(n+m)

	2. We can just copy nums2 to nums1 form m index. And using sort function to sort it.
	Time O((n+m)log(n+m)), space O(1)

	3.
	[1 2 3 0 0 0]
	[2 5 6]

	Instead of going from left to right, we can do from right to left.
	Choose larger number when comparing between nums1 and nums2. And fill it into right most.
	Keep continue to do it until nums1 or nums2 have no elements.
	If num2 still have elements left, we just override nums1 at the same index.
	Time: 0(n+m), space O(1)

*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for k := len(nums1) - 1; k >= 0; k-- {
		if i < 0 || j < 0 {
			break
		}

		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

	if j < 0 {
		return
	}

	for i := 0; i <= j; i++ {
		nums1[i] = nums2[i]
	}
}
