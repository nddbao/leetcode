package solution

/*
	leetcode: https://leetcode.com/problems/largest-component-size-by-common-factor/submissions/
*/

func largestComponentSize(nums []int) int {
	max := findMax(nums)
	ds := NewDisjointSet(max)
	for i := 0; i < len(nums); i++ {
		unionCommonFactor(ds, nums[i])
	}

	// count numbers that have the same root and choose max
	var result int
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		u := ds.Find(nums[i])
		freq[u]++
		result = getMax(result, freq[u])
	}

	return result
}

func findMax(nums []int) int {
	var max int
	for _, v := range nums {
		max = getMax(max, v)
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func unionCommonFactor(ds *DisjointSet, num int) {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			ds.Union(num, i)
			ds.Union(num, num/i)
		}
	}
}

func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	return &DisjointSet{parent: parent, rank: make([]int, n+1)}
}

type DisjointSet struct {
	parent []int
	rank   []int
}

func (d *DisjointSet) Find(v int) int {
	if d.parent[v] == v {
		return v
	}
	d.parent[v] = d.Find(d.parent[v])
	return d.parent[v]
}

func (d *DisjointSet) Union(u, v int) {
	u = d.Find(u)
	v = d.Find(v)

	if u == v {
		return
	}

	if d.rank[u] < d.rank[v] {
		u, v = v, u
	}
	d.parent[v] = u

	if d.rank[u] == d.rank[v] {
		d.rank[u]++
	}
}
