package solution

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var max int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := calculateArea(grid, i, j)
				if max < area {
					max = area
				}
			}
		}
	}
	return max
}

func calculateArea(grid [][]int, i, j int) int {
	if !isValidIndex(grid, i, j) || grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = -1
	result := 1
	for _, v := range connectedIndices {
		result += calculateArea(grid, i+v[0], j+v[1])
	}

	return result
}

func isValidIndex(grid [][]int, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

var connectedIndices = [][]int{
	[]int{-1, 0}, // top
	[]int{1, 0},  // bottom
	[]int{0, -1}, // left
	[]int{0, 1},  // right
}
