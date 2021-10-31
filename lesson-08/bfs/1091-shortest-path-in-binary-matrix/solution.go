package solution

/*
	leetcode: https://leetcode.com/problems/shortest-path-in-binary-matrix/
*/

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if n == 1 && grid[0][0] == 0 {
		return 1
	}

	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	return bfs(grid)
}

func bfs(grid [][]int) int {
	q := &Queue{}
	q.Enque([]int{0, 0})
	markVisited(grid, 0, 0)

	n := len(grid)
	level := 1
	for q.Size() > 0 {
		size := q.Size()
		level++
		for k := 0; k < size; k++ {
			node := q.Deque()
			for _, d := range directions {
				i, j := node[0]+d[0], node[1]+d[1]
				if validIdx(grid, i, j) && grid[i][j] == 0 {
					if i == n-1 && j == n-1 {
						return level
					}
					q.Enque([]int{i, j})
					markVisited(grid, i, j)
				}
			} // end loop
		} // end loop
	} // end loop
	return -1
}

func validIdx(grid [][]int, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid)
}

func markVisited(grid [][]int, i, j int) {
	grid[i][j] = 2
}

var directions = [][]int{
	[]int{-1, 0},  // top
	[]int{-1, -1}, // top-left
	[]int{-1, 1},  // top-right
	[]int{1, 0},   // bottom
	[]int{1, -1},  // bottom-left
	[]int{1, 1},   // bottom-right
	[]int{0, -1},  // left
	[]int{0, 1},   // right
}

type Queue struct {
	a [][]int
}

func (q *Queue) Enque(node []int) {
	q.a = append(q.a, node)
}

func (q *Queue) Deque() []int {
	node := q.a[0]
	q.a = q.a[1:]
	return node
}

func (q *Queue) Size() int {
	return len(q.a)
}
