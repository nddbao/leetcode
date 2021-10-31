package solution

func slidingPuzzle(board [][]int) int {
	start := getNode(board)
	end := &Node{s: "123450", emptyIdx: 5}
	if start.s == end.s {
		return 0
	}

	return bfs(start, end)
}

func bfs(start, end *Node) int {
	visited := make(map[string]struct{})
	visited[start.s] = struct{}{}

	q := &Queue{}
	q.Enque(start)

	level := 0
	for q.Size() > 0 {
		size := q.Size()
		level++
		for i1 := 0; i1 < size; i1++ {
			node := q.Deque()
			nextNodes := getNextNodes(node)
			for _, v := range nextNodes {
				if end.s == v.s {
					return level
				}

				_, exist := visited[v.s]
				if exist {
					continue
				}

				visited[v.s] = struct{}{}
				q.Enque(v)
			} // end loop
		} // end loop
	} // end loop
	return -1
}

func getNextNodes(node *Node) []*Node {
	var result []*Node
	for _, v := range adjacents[node.emptyIdx] {
		s := []byte(node.s)
		s[node.emptyIdx], s[v] = s[v], s[node.emptyIdx]
		result = append(result, &Node{s: string(s), emptyIdx: v})
	}
	return result
}

/*
   index in puzzle
   [[0 1 2]
    [3 4 5]]
*/
var adjacents = [][]int{
	{1, 3},    // 0
	{0, 2, 4}, // 1
	{1, 5},    // 2
	{0, 4},    // 3
	{1, 3, 5}, // 4
	{2, 4},    // 5
}

func getNode(board [][]int) *Node {
	var s []byte
	var emptyIdx int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			s = append(s, '0'+byte(board[i][j]))
			if board[i][j] == 0 {
				emptyIdx = i*3 + j
			}
		}
	}

	return &Node{s: string(s), emptyIdx: emptyIdx}
}

type Node struct {
	s        string
	emptyIdx int
}

type Queue struct {
	a []*Node
}

func (q *Queue) Enque(node *Node) {
	q.a = append(q.a, node)
}

func (q *Queue) Deque() *Node {
	node := q.a[0]
	q.a = q.a[1:]
	return node
}

func (q *Queue) Size() int {
	return len(q.a)
}
