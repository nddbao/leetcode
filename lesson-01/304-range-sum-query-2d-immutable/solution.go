package solution

/*
	leetcode: https://leetcode.com/problems/range-sum-query-2d-immutable/
*/
type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var upperLeft, upperRight, lowerLeft, lowerRight int

			if i-1 >= 0 && j-1 >= 0 {
				upperLeft = matrix[i-1][j-1]
			}

			if i-1 >= 0 {
				upperRight = matrix[i-1][j]
			}

			if j-1 >= 0 {
				lowerLeft = matrix[i][j-1]
			}

			lowerRight = matrix[i][j]

			matrix[i][j] = upperRight + lowerLeft - upperLeft + lowerRight
		}
	}
	return NumMatrix{preSum: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var x1, x2, x3 int
	if row1-1 >= 0 && col1-1 >= 0 {
		x1 = this.preSum[row1-1][col1-1]
	}

	if row1-1 >= 0 {
		x2 = this.preSum[row1-1][col2]
	}

	if col1-1 >= 0 {
		x3 = this.preSum[row2][col1-1]
	}

	return this.preSum[row2][col2] - x2 - x3 + x1
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
