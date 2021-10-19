package solution

/*
	leetcode: https://leetcode.com/problems/sum-of-unique-elements/
*/

/*
	We know range of value array will be [1..100]
	So we just go through, count value occurrence and put it in other array.
	And after that, we can sum up for value have occurrence equal to 1.

	Time complexity: O(n)
	Space complexity: O(1)

*/

func sumOfUnique(nums []int) int {
	var a [101]int
	for _, v := range nums {
		a[v]++
	}

	var s int
	for i, v := range a {
		if v == 1 {
			s += i
		}
	}

	return s
}
