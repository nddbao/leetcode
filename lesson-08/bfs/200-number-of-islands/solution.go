package solution

/*
	leetcode: https://leetcode.com/problems/number-of-islands/
*/

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				bfs(grid, i, j)
				count++
			}
		}
	}

	return count
}

func bfs(grid [][]byte, i, j int) {
	grid[i][j] = '2'
	queue := [][]int{[]int{i, j}}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		for _, v := range positionIndices {
			x, y := item[0]+v[0], item[1]+v[1]
			if isValidIndex(grid, x, y) && grid[x][y] == '1' {
				grid[x][y] = '2'
				queue = append(queue, []int{x, y})
			}
		}
	}
}

func isValidIndex(grid [][]byte, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

var positionIndices = [][]int{
	[]int{-1, 0}, // top
	[]int{1, 0},  // bottom
	[]int{0, -1}, // left
	[]int{0, 1},  // right
}
