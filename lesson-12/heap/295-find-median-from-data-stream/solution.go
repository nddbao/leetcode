package solution

import "container/heap"

/*
	leetcode: https://leetcode.com/problems/find-median-from-data-stream/
*/

/*
	We use to two heaps:
		+ max heap: stores smaller numbers
		+ min heap: stores larger numbers
	When a number is added to median finder,
	 => we have to make sure that len(maxHeap) - len(minHeap) in [0, 1]

	When findMedian:
		+ If len(maxHeap) - len(minHeap) = 0, we know that total number is even
			=> we have calculate average from peek of maxHeap and minHeap
		+ If len(maxHeap) - len(minHeap) = 1, we know that total number is old
			=> we only return peek of maxHeap

	Time complexity:
		AddNum: O(logN)
		FindMedian: O(1)
	Space complexity: O(N)

*/
type MedianFinder struct {
	maxHeap *MyHeap
	minHeap *MyHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: NewMaxHeap(),
		minHeap: NewMinHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 {
		heap.Push(this.maxHeap, num)
	} else {
		val := this.maxHeap.Peek().(int)
		if num <= val {
			heap.Push(this.maxHeap, num)
		} else {
			heap.Push(this.minHeap, num)
		}
	}

	if this.minHeap.Len() > this.maxHeap.Len() {
		val := heap.Pop(this.minHeap)
		heap.Push(this.maxHeap, val)
	} else if this.maxHeap.Len()-this.minHeap.Len() > 1 {
		val := heap.Pop(this.maxHeap)
		heap.Push(this.minHeap, val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	x1 := this.maxHeap.Peek().(int)
	if this.maxHeap.Len() == this.minHeap.Len() {
		x2 := this.minHeap.Peek().(int)
		return (float64(x1) + float64(x2)) / 2
	}

	return float64(x1)
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func NewMaxHeap() *MyHeap {
	return &MyHeap{
		compareFn: func(a, b int) bool { return a > b },
	}
}

func NewMinHeap() *MyHeap {
	return &MyHeap{
		compareFn: func(a, b int) bool { return a < b },
	}
}

type MyHeap struct {
	a         []int
	compareFn func(a, b int) bool
}

func (h *MyHeap) Len() int {
	return len(h.a)
}

func (h *MyHeap) Swap(i, j int) {
	h.a[i], h.a[j] = h.a[j], h.a[i]
}

func (h *MyHeap) Less(i, j int) bool {
	return h.compareFn(h.a[i], h.a[j])
}

func (h *MyHeap) Push(v interface{}) {
	x := v.(int)
	h.a = append(h.a, x)
}

func (h *MyHeap) Pop() interface{} {
	old := h.a
	n := len(old)
	x := old[n-1]
	h.a = old[:n-1]
	return x
}

func (h *MyHeap) Peek() interface{} {
	return h.a[0]
}
