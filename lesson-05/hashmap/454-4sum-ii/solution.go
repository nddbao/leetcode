package solution

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
