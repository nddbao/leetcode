package solution

/*
	leetcode: https://leetcode.com/problems/search-a-2d-matrix-ii/
*/

/*
	Row in ascending from left -> right
	Column in ascending from top -> bottom

    Let's  pick cell that starts at bottom-left (i=m-1, j=0)
 	There are 3 cases:
		+ current value == target: we found it.
		+ current value < target:  check right cell of current cell.
		+ current value > target: check up cell of current cell

	We continue until we out of index and we found value == target.


	Time complexity: O(m + n)
	Space complexity: O(1)
*/

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	i, j := m-1, 0 // bottom-left
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
