package solution

/*
	leetcode: https://leetcode.com/problems/path-with-minimum-effort/
*/

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	if m == 1 && n == 1 {
		return 0
	}

	var result int
	l, h := 0, 1000*1000
	for l < h {
		mid := l + (h-l)/2
		visited := initVisited(m, n)
		if dfs(heights, visited, mid, 0, 0) {
			result = mid
			h = mid
		} else {
			l = mid + 1
		}
	}

	return result
}

func dfs(heights [][]int, visited [][]bool, target, i, j int) bool {
	visited[i][j] = true

	if isBottomRightCell(heights, i, j) {
		return true
	}

	// move next to 4 directions
	for _, d := range directions {
		x, y := i+d[0], j+d[1]
		if !isValidIdx(heights, x, y) || visited[x][y] {
			continue
		}

		if abs(heights[x][y]-heights[i][j]) > target {
			continue
		}

		if dfs(heights, visited, target, x, y) {
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

func isBottomRightCell(heights [][]int, i, j int) bool {
	return i == len(heights)-1 && j == len(heights[0])-1
}

func isValidIdx(heights [][]int, i, j int) bool {
	return i >= 0 && i < len(heights) && j >= 0 && j < len(heights[0])
}

func initVisited(m, n int) [][]bool {
	a := make([][]bool, m)
	for i := 0; i < len(a); i++ {
		a[i] = make([]bool, n)
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
