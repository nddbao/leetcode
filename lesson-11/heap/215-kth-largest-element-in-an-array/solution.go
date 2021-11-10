package solution

import "container/heap"

/*
	leetcode: https://leetcode.com/problems/kth-largest-element-in-an-array/
*/

/*
	We use init a max heap based on current array.
	Then we pop out element of heap for K times.
	Last item is our result.

	Time complexity: O(n + klogn)
		+ init heap: O(n)
		+ get k-th largest: O(klogN)

    Space complexity: O(1)
*/
func findKthLargest(nums []int, k int) int {
	h := NewIntHeap(nums)
	heap.Init(h)
	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	result := heap.Pop(h).(int)

	return result
}

func NewIntHeap(nums []int) *IntHeap {
	h := IntHeap(nums)
	return &h
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
