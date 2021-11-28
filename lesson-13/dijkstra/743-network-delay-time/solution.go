package solution

import (
	"container/heap"
	"math"
)

/*
	leetcode: https://leetcode.com/problems/network-delay-time/
*/

/*
	We use dijkstra to solve this problem.
	First, we go from node K and update cost for nodes that node K can connected to.
	Then we choose next node that has min cost and do the same as node K.
	We will stop when we cannot choose any.

	Time complexity:
		NewNodeTable: O(N + E) whether N is number of node, E is number of edges
		Loop: O(N^2)
			+ Outside: O(N)
			+ GetMinNode: O(ElogE)
			+ Inside: O(N)

	Space complexity: O(N + E)

*/
func networkDelayTime(times [][]int, n int, k int) int {
	table := NewNodeTable(times, n, k)
	h := NewMinHeap(&Node{Val: k, Cost: 0})
	max := 0
	count := 0
	for {
		// get node that has min cost and mark visited
		node := GetMinNode(table, h)
		if node == nil {
			break
		}
		src := node.Val
		table.MarkVisited(src)

		// update max travel time and count number of node
		max = GetMax(max, node.Cost)
		count++

		// update cost for other nodes
		edges := table.GetDirectedEdges(node.Val)
		if len(edges) == 0 {
			continue
		}

		for _, v := range edges {
			e := Edge(v)
			target := e.Target()
			if table.IsVisited(target) {
				continue
			}

			cost := table.GetCost(src) + e.Cost()
			if cost < table.GetCost(target) {
				table.UpdateCost(target, cost)
				heap.Push(h, &Node{Val: target, Cost: cost})
			}
		} // end loop

	} // end loop

	if count == n {
		return max
	}

	return -1
}

func GetMinNode(table *NodeTable, h *MinHeap) *Node {
	if h.Len() == 0 {
		return nil
	}
	node := heap.Pop(h).(*Node)
	for table.IsVisited(node.Val) {
		if h.Len() == 0 {
			return nil
		}
		node = heap.Pop(h).(*Node)
	}

	return node
}

func NewNodeTable(times [][]int, n, k int) *NodeTable {
	return &NodeTable{
		cost:     NewCost(n, k),
		visited:  make([]bool, n+1),
		edgeList: NewDirectedEdgesList(times)}
}

func NewCost(n, k int) []int {
	a := make([]int, n+1)
	for i := 0; i < len(a); i++ {
		a[i] = math.MaxInt32
	}
	a[k] = 0
	return a
}

func NewDirectedEdgesList(times [][]int) map[int][][]int {
	list := make(map[int][][]int)
	for _, v := range times {
		src := v[0]
		list[src] = append(list[src], v)
	}
	return list
}

type NodeTable struct {
	cost     []int
	visited  []bool
	edgeList map[int][][]int
}

func (c *NodeTable) GetCost(i int) int {
	return c.cost[i]
}

func (c *NodeTable) UpdateCost(i, val int) {
	c.cost[i] = val
}

func (c *NodeTable) MarkVisited(i int) {
	c.visited[i] = true
}

func (c *NodeTable) IsVisited(i int) bool {
	return c.visited[i]
}

func (c *NodeTable) GetDirectedEdges(i int) [][]int {
	return c.edgeList[i]
}

type Edge []int

func (e Edge) Source() int {
	return e[0]
}

func (e Edge) Target() int {
	return e[1]
}

func (e Edge) Cost() int {
	return e[2]
}

type Node struct {
	Val  int
	Cost int
}

func NewMinHeap(node *Node) *MinHeap {
	return &MinHeap{a: []*Node{node}}
}

type MinHeap struct {
	a []*Node
}

func (h MinHeap) Len() int {
	return len(h.a)
}

func (h MinHeap) Swap(i, j int) {
	h.a[i], h.a[j] = h.a[j], h.a[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h.a[i].Cost < h.a[j].Cost
}

func (h *MinHeap) Push(v interface{}) {
	h.a = append(h.a, v.(*Node))
}

func (h *MinHeap) Pop() interface{} {
	old := h.a
	n := len(old)
	x := old[n-1]
	h.a = old[:n-1]
	return x
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
