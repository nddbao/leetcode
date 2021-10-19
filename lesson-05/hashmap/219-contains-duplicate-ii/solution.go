package solution

/*
	leetcode: https://leetcode.com/problems/contains-duplicate-ii/
*/

/*
	We can use brute force check combination of two values in the array.
	But time complexity will be O(n^2)

	So we can do it by one pass through the array by using map to store (value, index)
	Every iteration, we check current value existing in map.
	If existing, check two indexes satisfy  our requirement or not.

	Time complexity: O(n)
	Space complexity: O(n)
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
