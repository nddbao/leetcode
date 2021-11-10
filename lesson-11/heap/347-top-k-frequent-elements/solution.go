package solution

import "container/heap"

/*
	leetcode: https://leetcode.com/problems/top-k-frequent-elements/
*/

func topKFrequent(nums []int, k int) []int {
	distincts := make(map[int]int)
	for _, v := range nums {
		distincts[v]++
	}

	a := make([][2]int, 0)
	h := NewHeap(a)
	for key, value := range distincts {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		v := heap.Pop(h).([2]int)
		result[i] = v[0]
	}
	return result
}

func NewHeap(a [][2]int) *ItemHeap {
	h := ItemHeap(a)
	return &h
}

type ItemHeap [][2]int

func (h ItemHeap) Len() int {
	return len(h)
}

func (h ItemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h ItemHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h *ItemHeap) Push(v interface{}) {
	*h = append(*h, v.([2]int))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}
