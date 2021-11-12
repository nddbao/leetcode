package solution

/*
	leetcode: https://leetcode.com/problems/search-a-2d-matrix-ii/
*/

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	i, j := m-1, 0
	for isValidIdx(m, n, i, j) {
		val := matrix[i][j]
		if val == target {
			return true
		}

		d := getDirection(val, target)
		i, j = i+d[0], j+d[1]
	}

	return false
}

var (
	up    = []int{-1, 0}
	right = []int{0, 1}
)

func getDirection(val, target int) []int {
	if val < target {
		return right
	}

	return up
}

func isValidIdx(rows, columns, i, j int) bool {
	return i >= 0 && i < rows && j >= 0 && j < columns
}
