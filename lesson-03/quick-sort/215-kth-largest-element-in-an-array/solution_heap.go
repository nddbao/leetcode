package solution

import "container/heap"

/*
	leetcode: https://leetcode.com/problems/kth-largest-element-in-an-array/
*/

func findKthLargestHeap(nums []int, k int) int {
	var h = MyHeap([]int{})
	heap.Init(&h)

	for _, v := range nums {
		heap.Push(&h, v)
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	return heap.Pop(&h).(int)
}

type MyHeap []int

func (m MyHeap) Len() int {
	return len(m)
}

func (m MyHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MyHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m *MyHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}

func (m *MyHeap) Pop() interface{} {
	v := (*m)[(*m).Len()-1]
	*m = (*m)[:(*m).Len()-1]
	return v
}
