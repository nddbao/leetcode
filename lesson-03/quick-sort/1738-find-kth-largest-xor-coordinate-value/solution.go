package solution

/*
	leetcode: https://leetcode.com/problems/find-kth-largest-xor-coordinate-value/
*/

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])

	// calculate XOR value
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = XORValue(matrix, i, j)
		}
	}

	// sort matrix
	a := Arr(matrix)
	sort.Sort(a)

	x, y := a.getCoordinate(k - 1)
	return matrix[x][y]
}

func XORValue(matrix [][]int, i, j int) int {
	result := matrix[i][j]

	if i-1 >= 0 && j-1 >= 0 {
		result = result ^ matrix[i-1][j-1]
	}

	if i-1 >= 0 {
		result = result ^ matrix[i-1][j]
	}

	if j-1 >= 0 {
		result = result ^ matrix[i][j-1]
	}

	return result
}

type Arr [][]int

func (a Arr) Len() int {
	return len(a) * len(a[0])
}

func (a Arr) Less(i, j int) bool {
	x1, y1 := a.getCoordinate(i)
	x2, y2 := a.getCoordinate(j)
	return a[x1][y1] > a[x2][y2]
}

func (a Arr) Swap(i, j int) {
	x1, y1 := a.getCoordinate(i)
	x2, y2 := a.getCoordinate(j)
	a[x1][y1], a[x2][y2] = a[x2][y2], a[x1][y1]
}

func (a Arr) getCoordinate(x int) (int, int) {
	return x / len(a[0]), x % len(a[0])
}
