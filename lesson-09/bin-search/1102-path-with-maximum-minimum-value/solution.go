package solution

/*
	leetcode: https://leetcode.com/problems/path-with-maximum-minimum-value/
*/

/*
	Note: remember to check grid[0][0] >= mid
*/
func maximumMinimumPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	l, h := 0, int(10e9)
	result := 0
	for l < h {
		mid := l + (h-l)/2
		visited := initVisited(m, n)
		if grid[0][0] >= mid && dfs(grid, visited, mid, 0, 0) {
			result = mid
			l = mid + 1
		} else {
			h = mid
		}
	}
	return result
}

func dfs(grid [][]int, visited [][]bool, target, i, j int) bool {
	visited[i][j] = true

	if isEnd(grid, i, j) {
		return true
	}

	for _, d := range directions {
		x, y := i+d[0], j+d[1]

		if !isValidIdx(grid, x, y) || visited[x][y] {
			continue
		}

		if grid[x][y] >= target && dfs(grid, visited, target, x, y) {
			return true
		}
	}
	return false
}

var directions = [][]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

func isEnd(grid [][]int, i, j int) bool {
	return i == len(grid)-1 && j == len(grid[0])-1
}

func isValidIdx(grid [][]int, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func initVisited(m, n int) [][]bool {
	a := make([][]bool, m)
	for i := 0; i < m; i++ {
		a[i] = make([]bool, n)
	}
	return a
}
