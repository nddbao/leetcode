package solution

/*
	leetcode: https://leetcode.com/problems/two-sum
*/

/*
	Basically, we can use brute force to solve this problem.
	By checking every combination of two values whether they can sum up to target.
	But time complexity will be O(n^2).

	We can do this better with one pass by using map.
	Iterate through the array and map value to index.
	we check (target- value) that is exist in map.
	If exist, we know the result to return.

	Time complexity: O(n)
	Space complexity: O(n)
*/

func twoSum(nums []int, target int) []int {
	distinct := make(map[int]int)
	for i1, v := range nums {
		i2, exist := distinct[target-v]
		if exist {
			return []int{i1, i2}
		}

		distinct[v] = i1
	}

	return nil
}
