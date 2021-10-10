package solution

/*
	leetcode: https://leetcode.com/problems/merge-sorted-array/
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
