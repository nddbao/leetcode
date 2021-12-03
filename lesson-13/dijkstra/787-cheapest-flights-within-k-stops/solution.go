package solution

import (
	"container/heap"
	"math"
)

/*
	leetcode: https://leetcode.com/problems/cheapest-flights-within-k-stops/
*/

/*
	We use dijkstra to solve this.
	We go from src to check connected flights.
	Then we calculate the cost as well as adding (city, stop -1, cost) to min heap.
	We will not add to min heap with some conditions:
		+ stop <= 0
		+ same city: ( condition: comparing by stop and cost) new one have higher value than existing one in min heap.

	When we reach destination, we can update our min.

	Time complexity: k*M*log(M*k)
		buildCosts: O(N) whether N is number of cities
		buildAdjacentList: O(M) whether M is number of flights
		loop to find min: k*M*log(M*k)

	Space complexity: O(M*k + N)

*/

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	fd := NewFacade(n, flights, src, k)
	min := math.MaxInt32
	for fd.HasNextItem() {
		item := fd.NextItem()
		connectedFlights := fd.FlightsFrom(item.City)

		if item.Stop == 0 {
			continue
		}

		for _, v := range connectedFlights {
			city := v.End()
			cost := item.Cost + v.Price()

			if cost > min {
				continue
			}

			if city == dst {
				min = getMin(min, cost)
			} else {
				newItem := NewItem(city, cost, item.Stop-1)
				fd.PushItem(newItem)
			}
		} // end lop

	} // end loop

	if min == math.MaxInt32 {
		return -1
	}

	return min
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func NewFacade(n int, flights [][]int, src int, k int) Facade {
	list := buildAdjacentList(flights)
	srcItem := NewItem(src, 0, k+1)
	h := &MinHeap{}
	heap.Push(h, srcItem)
	costs := buildCosts(n, src, k)
	return Facade{adjList: list, minHeap: h, minCosts: costs}
}

type Facade struct {
	adjList  map[int][]Flight
	minHeap  *MinHeap
	minCosts []*Item
}

func (f Facade) HasNextItem() bool {
	return f.minHeap.Len() > 0
}

func (f Facade) NextItem() *Item {
	return heap.Pop(f.minHeap).(*Item)
}

func (f Facade) FlightsFrom(city int) []Flight {
	return f.adjList[city]
}

func (f Facade) PushItem(item *Item) {
	old := f.minCosts[item.City]

	if !compareItemLessThan(item, old) {
		return
	}

	heap.Push(f.minHeap, item)
	f.minCosts[item.City] = item
}

func compareItemLessThan(a, b *Item) bool {
	if a.Stop > b.Stop {
		return true
	}

	if a.Stop == b.Stop {
		return a.Cost < b.Cost
	}

	if a.Cost < b.Cost {
		return true
	}
	return false
}

func buildCosts(n, src, k int) []*Item {
	costs := make([]*Item, n)
	for i := 0; i < len(costs); i++ {
		costs[i] = NewItem(i, math.MaxInt32, -1)
	}
	costs[src] = NewItem(src, 0, k+1)
	return costs
}

func buildAdjacentList(flights [][]int) map[int][]Flight {
	list := make(map[int][]Flight)
	for _, v := range flights {
		f := Flight(v)
		start := f.Start()
		list[start] = append(list[start], f)
	}
	return list
}

type Flight []int

func (f Flight) Start() int {
	return f[0]
}

func (f Flight) End() int {
	return f[1]
}

func (f Flight) Price() int {
	return f[2]
}

func NewItem(city, cost, stop int) *Item {
	return &Item{City: city, Cost: cost, Stop: stop}
}

type Item struct {
	City int
	Stop int
	Cost int
}

type MinHeap struct {
	a []*Item
}

func (h MinHeap) Len() int {
	return len(h.a)
}

func (h MinHeap) Less(i, j int) bool {
	return h.a[i].Cost < h.a[j].Cost
}

func (h MinHeap) Swap(i, j int) {
	h.a[i], h.a[j] = h.a[j], h.a[i]
}

func (h *MinHeap) Push(v interface{}) {
	h.a = append(h.a, v.(*Item))
}

func (h *MinHeap) Pop() interface{} {
	old := h.a
	n := len(old)
	x := old[n-1]
	h.a = old[:n-1]
	return x
}
