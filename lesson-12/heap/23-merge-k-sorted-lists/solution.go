package solution

import "container/heap"

/*
	leetcode: https://leetcode.com/problems/merge-k-sorted-lists/
*/

/*
		We use min heap and build from first element of each list node.
	We are sure that we will always have smallest element in heap.

		Using a loop to get smallest element x in heap and add to result list.
    On the other hand, we check next element of x and push to heap.

	Time complexity: O(NlogK) whether N is total nodes, K is len of array ListNode
		buildHeap: O(K)
		build sort list: O(NlogK)

	Space complexity: O(N + K)
*/

func mergeKLists(lists []*ListNode) *ListNode {
	dump := &ListNode{}
	curr := dump

	h := buildHeap(lists)
	for h.Len() > 0 {
		x := heap.Pop(h).(*ListNode)

		// add next item to min heap
		next := x.Next
		if next != nil {
			heap.Push(h, next)
		}

		// set value for next current
		curr.Next = x
		curr = curr.Next
	}

	return dump.Next
}

func buildHeap(lists []*ListNode) *MinHeap {
	var a []*ListNode
	for _, v := range lists {
		if v != nil {
			a = append(a, v)
		}
	}

	h := MinHeap(a)
	heap.Init(&h)
	return &h
}

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *MinHeap) Push(v interface{}) {
	x := v.(*ListNode)
	*h = append(*h, x)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type ListNode struct {
	Val  int
	Next *ListNode
}
