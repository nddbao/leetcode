package solution

/*
	leetcode: https://leetcode.com/problems/find-anagram-mappings/
*/

func anagramMappings(nums1 []int, nums2 []int) []int {
	distinct := make(map[int][]int)
	for i, v := range nums2 {
		distinct[v] = append(distinct[v], i)
	}

	var result []int
	for _, v := range nums1 {
		a := distinct[v]
		result = append(result, a[len(a)-1])
		distinct[v] = a[:len(a)-1]
	}

	return result
}
