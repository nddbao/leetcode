package solution

/*
	leetcode: https://leetcode.com/problems/minimum-knight-moves/
*/

func minKnightMoves(x int, y int) int {
	if x == 0 && y == 0 {
		return 0
	}

	start := []int{300, 300}
	end := []int{x + 300, y + 300}
	return bfs(start, end)
}

func bfs(start, end []int) int {
	visited := make([]bool, 602*602)
	q := &Queue{}
	q.Enque(start)
	markVisited(visited, start[0], start[1])

	level := 0
	for q.Size() > 0 {
		size := q.Size()
		level++
		for i1 := 0; i1 < size; i1++ {
			node := q.Deque()
			for _, d := range directions {
				i, j := node[0]+d[0], node[1]+d[1]

				if !isValidIdx(i, j) {
					continue
				}

				if i == end[0] && j == end[1] {
					return level
				}

				if isVisited(visited, i, j) {
					continue
				}

				q.Enque([]int{i, j})
				markVisited(visited, i, j)
			} // end loop
		} // end loop
	} // end loop
	return -1
}

func isValidIdx(i, j int) bool {
	return i >= 0 && i <= 601 && j >= 0 && j <= 601
}
func markVisited(visited []bool, i, j int) {
	visited[i*602+j] = true
}

func isVisited(visited []bool, i, j int) bool {
	return visited[i*602+j]
}

var directions = [][]int{
	{-2, -1}, // top-left
	{-2, 1},  // top-right
	{2, -1},  // bottom-left
	{2, 1},   // bottom-right
	{-1, -2}, // left-top
	{1, -2},  // left-bottom
	{-1, 2},  // right-top
	{1, 2},   // right-bottom
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
