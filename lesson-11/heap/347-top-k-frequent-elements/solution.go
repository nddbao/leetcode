package solution

import "container/heap"

/*
	leetcode: https://leetcode.com/problems/top-k-frequent-elements/
*/

/*
	We need to know frequency of each number.
	So we to check all of elements in array. Then storing them to a map.

	We go through this map and build a min heap have at maximum at k elements.
	Every iteration, we add an item{key=index, value=freq} to min heap.
	If len(heap) > k, we just pop an item out of heap.

	Finally, we have k largest frequency items.

	Time complexity: O(nlogk)
		+ countFrequencies -> O(n)
		+ buildMinHeap -> O(nlogk)
		+ get resutl: O(klogk)

	Space complexity: O(n+k)
*/
func topKFrequent(nums []int, k int) []int {
	freqs := countFrequencies(nums)
	h := buildMinHeap(freqs, k)

	result := make([]int, k)
	for i := 0; i < k; i++ {
		v := heap.Pop(h).([2]int)
		result[i] = v[0]
	}
	return result
}

func buildMinHeap(freqs map[int]int, k int) *ItemHeap {
	a := make([][2]int, 0)
	h := NewHeap(a)
	for key, value := range freqs {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return h
}

func countFrequencies(nums []int) map[int]int {
	freqs := make(map[int]int)
	for _, v := range nums {
		freqs[v]++
	}
	return freqs
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
