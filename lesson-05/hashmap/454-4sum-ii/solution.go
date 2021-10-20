package solution

/*
	leetcode: https://leetcode.com/problems/4sum-ii/
*/

/*
	Using an map to store frequent sum from all combination
	that can sum up from nums1 and nums1.

	Then we check sum of all combination that can sum up from num3 and nums4.
	We have: s = nums3 + num4 = - (nums1 + nums2)
	if  (-s) found in map, we know for sure: nums1 + nums2 + nums3 + nums4 = 0
	Just count it.

	Time complexity: O(n^2)
	Space complexity: O(n^2)

*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mSum := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			mSum[v1+v2]++
		}
	}

	count := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			s := v3 + v4
			v, ok := mSum[-s]
			if ok {
				count += v
			}
		}
	}

	return count
}
