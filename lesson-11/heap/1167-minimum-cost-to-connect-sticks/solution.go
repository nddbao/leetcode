package solution

import "container/heap"

/*
	leetcode: https://leetcode.com/problems/minimum-cost-to-connect-sticks/
*/

/*
	If we want to have minimum cost, each time we should connect two sticks that have minimum length.
	So we will min Heap in this problems. First, we init Heap based on input array.
	Then we do loop: pop 2 items from heap, calculate it and push result back to heap
	Until we have only one element left in heap.

	Time complexity: O(n) + O(logn) => O(n)
	Space complexity: O(1)
*/
func connectSticks(sticks []int) int {
	h := NewHeap(sticks)

	var cost int
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)

		c := a + b
		cost += c
		heap.Push(h, c)
	}

	return cost
}

func NewHeap(a []int) *IntHeap {
	h := IntHeap(a)
	heap.Init(&h)
	return &h
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
