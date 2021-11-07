package solution

/*
	leetcode: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
*/

/*
	Using two pointers left and right:
		+ Left: start from beginning
		+ Right: start from the end
	Loop through the array and compute sum = numbers[left] + numbers[right]
	 + if sum < target, we need increase our sum by moving left one more ahead.
	 + if sum > target, we need decrease our sum by moving right one more back.
	 + if sum = target, it means we find out the answer. Just return [left + 1, right + 1] (1-indexed)

	Time Complexity: O(n)
	Space Complexity: O(1)

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
