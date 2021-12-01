package solution

/*
	leetcode: https://leetcode.com/problems/graph-connectivity-with-threshold/
*/

func areConnected(n int, threshold int, queries [][]int) []bool {
	ds := NewDisjointSet(n + 1)
	for i := threshold + 1; i <= n; i++ {
		for j := i + i; j <= n; j += i {
			ds.Union(i, j)
		} // end loop
	} // end loop

	result := make([]bool, len(queries))
	for idx, q := range queries {
		u := ds.Find(q[0])
		v := ds.Find(q[1])
		if u == v {
			result[idx] = true
		}
	} // end loop

	return result
}

func NewDisjointSet(n int) *DisjointSet {
	a := make([]int, n)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	return &DisjointSet{parent: a, rank: make([]int, n)}
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

	// update parent
	if d.rank[u] < d.rank[v] {
		u, v = v, u
	}
	d.parent[v] = u

	// update rank
	if d.rank[u] == d.rank[v] {
		d.rank[u]++
	}
}
