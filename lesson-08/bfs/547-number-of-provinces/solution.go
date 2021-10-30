package solution

/*
	leetcode: https://leetcode.com/problems/number-of-provinces/
*/

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	if n <= 1 {
		return n
	}

	var count int
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs(isConnected, visited, i)
			count++
		}
	}

	return count
}

func bfs(isConnected [][]int, visited []bool, idx int) {
	q := &Queue{}
	q.Enque(idx)
	visited[idx] = true
	for q.Size() > 0 {
		node := q.Deque()
		for i := 0; i < len(isConnected); i++ {
			if isConnected[node][i] == 1 && !visited[i] {
				visited[i] = true
				q.Enque(i)
			}
		}
	}
}

type Queue struct {
	a []int
}

func (q *Queue) Enque(node int) {
	q.a = append(q.a, node)
}

func (q *Queue) Deque() int {
	node := q.a[0]
	q.a = q.a[1:]
	return node
}

func (q *Queue) Size() int {
	return len(q.a)
}
