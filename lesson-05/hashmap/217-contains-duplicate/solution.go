package solution

/*
	leetcode: https://leetcode.com/problems/contains-duplicate/
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
