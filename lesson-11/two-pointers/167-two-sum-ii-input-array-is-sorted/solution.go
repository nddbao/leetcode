package solution

/*
	leetcode: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
*/

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if target < sum {
			r--
		} else {
			l++
		}
	}
	return nil
}
