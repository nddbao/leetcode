package solution

/*
	leetcode: https://leetcode.com/problems/the-k-strongest-values-in-an-array/
*/

import "sort"

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	m := arr[(len(arr)-1)/2]

	sort.Slice(arr, func(i, j int) bool {
		a := abs(arr[i] - m)
		b := abs(arr[j] - m)
		return a > b || (a == b && arr[i] > arr[j])
	})

	return arr[:k]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
