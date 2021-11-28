package solution

import "container/heap"

/*
	leetcode: https://leetcode.com/problems/path-with-maximum-probability/
*/

/*
	We use dijkstra to solve this problem.

	We go from start and check all connected node.
	After that we update cost if that cost > current cost of this node.

	We pick next largest one and continue
	until we reach end node or we have visited all of nodes.


	Time complexity: O(N*E*LogE)
		NewCost: O(N)E
		NewEdges: O(E)
		Outside loop: O(N)
		PopNodeWithMaxProbability: O(logE)
		Inside loop: O(E)
		UpdateCostL O(logE)
	Space complexity: O(N + E)
*/

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	table := NewTableNode(n, edges, succProb, start)
	for {
		// get node that has max probability
		node, ok := table.PopNodeWithMaxProbability()
		if !ok {
			break
		}

		// check if we reach destination
		if node == end {
			return table.GetCost(node)
		}
		table.MarkVisited(node)

		// get connected nodes and upate cost
		edges := table.GetEdges(node)
		for _, v := range edges {
			e := Edge(v)
			dest := e.GetDestination()
			if table.IsVisited(dest) {
				continue
			}

			cost := table.GetCost(node) * e.Cost()
			if cost > table.GetCost(dest) {
				table.UpdateCost(dest, cost)
			}
		} // end loop
	} // end loop

	return 0
}

func NewTableNode(n int, edges [][]int, succProb []float64, start int) *TableNode {
	return &TableNode{
		cost:    NewCost(n, start),
		edges:   NewEdges(edges, succProb, n),
		visited: make([]bool, n),
		maxHeap: NewMaxHeap(&Item{Node: start, Cost: 0}),
	}
}

func NewEdges(edges [][]int, succProb []float64, n int) [][][]interface{} {
	m := make([][][]interface{}, n)
	for i := 0; i < len(edges); i++ {
		e := edges[i]
		e1 := []interface{}{e[1], succProb[i]}
		e2 := []interface{}{e[0], succProb[i]}

		k1 := e[0]
		m[k1] = append(m[k1], e1)

		k2 := e[1]
		m[k2] = append(m[k2], e2)
	}
	return m
}

func NewCost(n, start int) []float64 {
	cost := make([]float64, n)
	cost[start] = 1
	return cost
}

type TableNode struct {
	cost    []float64
	visited []bool
	edges   [][][]interface{}
	maxHeap *MaxHeap
}

func (t *TableNode) MarkVisited(node int) {
	t.visited[node] = true
}

func (t *TableNode) IsVisited(node int) bool {
	return t.visited[node]
}

func (t *TableNode) GetCost(node int) float64 {
	return t.cost[node]
}

func (t *TableNode) UpdateCost(node int, val float64) {
	t.cost[node] = val
	heap.Push(t.maxHeap, &Item{Node: node, Cost: val})
}

func (t *TableNode) GetEdges(node int) [][]interface{} {
	return t.edges[node]
}

func (t *TableNode) PopNodeWithMaxProbability() (int, bool) {
	h := t.maxHeap
	for h.Len() > 0 {
		item := heap.Pop(h).(*Item)
		if !t.IsVisited(item.Node) {
			return item.Node, true
		}
	}

	return 0, false
}

type Item struct {
	Node int
	Cost float64
}

type Edge []interface{}

func (e Edge) Cost() float64 {
	return e[1].(float64)
}

func (e Edge) GetDestination() int {
	return e[0].(int)
}

func NewMaxHeap(item *Item) *MaxHeap {
	return &MaxHeap{a: []*Item{item}}
}

type MaxHeap struct {
	a []*Item
}

func (h MaxHeap) Len() int {
	return len(h.a)
}

func (h MaxHeap) Swap(i, j int) {
	h.a[i], h.a[j] = h.a[j], h.a[i]
}

func (h MaxHeap) Less(i, j int) bool {
	return h.a[i].Cost > h.a[j].Cost
}

func (h *MaxHeap) Push(v interface{}) {
	h.a = append(h.a, v.(*Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := h.a
	n := len(old)
	x := old[n-1]
	h.a = old[:n-1]
	return x
}
