package sollution

import "container/heap"

/*
	leetcode: https://leetcode.com/problems/kth-largest-element-in-a-stream/
*/

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	h := IntHeap(nums)
	heap.Init(&h)
	for h.Len() > k {
		heap.Pop(&h)
	}
	return KthLargest{h: &h, k: k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)
	for this.h.Len() > this.k {
		heap.Pop(this.h)
	}

	x := heap.Pop(this.h)
	heap.Push(this.h, x)
	return x.(int)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

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
