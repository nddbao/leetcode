package solution

/*
	leetcode: https://leetcode.com/problems/sum-of-unique-elements/
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
