package solution

/*
	leetcode: https://leetcode.com/problems/contains-duplicate/
*/

/*
	First thinking, we can use brute force to check every element in array.
	So time complexity will be O(n^2)

	Let's try to use map. We loop through the array and put element in map.
	If we see current element already exists in map, we know that is duplicate.

	Time complexity: O(n)
	Space complexity: O(n)

*/
func containsDuplicate(nums []int) bool {
	distinct := make(map[int]struct{})
	for _, v := range nums {
		_, exist := distinct[v]
		if exist {
			return true
		}
		distinct[v] = struct{}{}
	}
	return false
}
