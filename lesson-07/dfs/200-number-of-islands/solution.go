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
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if !isValidIndex(grid, i, j) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'
	for _, v := range connectedIndices {
		dfs(grid, i+v[0], j+v[1])
	}
}

func isValidIndex(grid [][]byte, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

var connectedIndices = [][]int{
	[]int{-1, 0}, // top
	[]int{1, 0},  // bottom
	[]int{0, -1}, // left
	[]int{0, 1},  // right
}
