package solution

/*
	leetcode: https://leetcode.com/problems/contains-duplicate-ii/
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	distinctElements := make(map[int]int)
	for j := 0; j < len(nums); j++ {
		i, exist := distinctElements[nums[j]]
		if exist && j-i <= k {
			return true
		}

		distinctElements[nums[j]] = j
	}

	return false
}
