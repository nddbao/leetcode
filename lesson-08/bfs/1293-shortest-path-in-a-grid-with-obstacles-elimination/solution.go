package solution

/*
	leetcode: https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
*/

func shortestPath(grid [][]int, k int) int {
	if len(grid) == 1 && len(grid[0]) == 1 {
		return 0
	}

	return bfs(grid, k)
}

func bfs(grid [][]int, k int) int {
	q := &Queue{}
	q.Enque([]int{0, 0, 0})

	m, n := len(grid), len(grid[0])

	// init graph to track all visited nodes
	g := &Graph{
		store: make([]map[int]struct{}, m*n),
		n:     n,
	}
	for i := 0; i <= k; i++ {
		g.markVisited(0, 0, i)
	}

	level := 0
	for q.Size() > 0 {
		size := q.Size()
		level++
		for i1 := 0; i1 < size; i1++ {
			node := q.Deque()
			for _, d := range directions {
				i, j := node[0]+d[0], node[1]+d[1]

				if !isValidIdx(grid, i, j) {
					continue
				}

				// reach bottom-right of matrix
				if i == m-1 && j == n-1 {
					return level
				}

				// check number elimination
				numberOfElimination := node[2]
				if grid[i][j] != 0 {
					if numberOfElimination >= k {
						continue
					}
					numberOfElimination++
				}

				if g.isVisited(i, j, numberOfElimination) {
					continue
				}

				q.Enque([]int{i, j, numberOfElimination})
				g.markVisited(i, j, numberOfElimination)
			} // end loop
		} // end loop
	} // end loop
	return -1
}

func isValidIdx(grid [][]int, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

var directions = [][]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

type Graph struct {
	store []map[int]struct{}
	n     int
}

func (g *Graph) markVisited(i, j, k int) {
	idx := i*g.n + j

	val := g.store[idx]
	if val == nil {
		val = make(map[int]struct{})
	}

	val[k] = struct{}{}
	g.store[idx] = val
}

func (g *Graph) isVisited(i, j, k int) bool {
	idx := i*g.n + j
	val := g.store[idx]
	if val == nil {
		return false
	}

	_, exist := val[k]

	return exist
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
