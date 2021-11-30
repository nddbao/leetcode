package solution

/*
	leetcode: https://leetcode.com/problems/satisfiability-of-equality-equations/
*/

/*
	We use union-find to merge all variable that have same value to same set.
	Then we check variable that have different value:
		+ same set: not correct -> return false.
		+ not same: correct, continue to check

	Time complexity: O(N)
	Space complexity: O(26) -> O(1)
*/
func equationsPossible(equations []string) bool {
	ds := NewDisjointSet()

	// merge all variables that have same value
	for _, s := range equations {
		eq := Equation(s)
		if !eq.IsEqual() {
			continue
		}

		first, second := eq.FirstVal(), eq.SecondVal()
		ds.Union(first, second)

	} // end loop

	// check variables that have different value
	for _, s := range equations {
		eq := Equation(s)
		if eq.IsEqual() {
			continue
		}

		first, second := eq.FirstVal(), eq.SecondVal()
		u, v := ds.Find(first), ds.Find(second)
		if u == v {
			return false
		}
	} // end loop

	return true
}

type Equation string

func (e Equation) FirstVal() byte {
	return e[0] - 'a'
}

func (e Equation) SecondVal() byte {
	return e[3] - 'a'
}

func (e Equation) IsEqual() bool {
	return e[1] == '='
}

func NewDisjointSet() *DisjointSet {
	n := byte('z' - 'a' + 1)
	parent := make([]byte, n)
	rank := make([]int, n)
	for i := 0; i < len(parent); i++ {
		parent[i] = byte(i)
	}

	return &DisjointSet{parent: parent, rank: rank}
}

type DisjointSet struct {
	parent []byte
	rank   []int
}

func (d *DisjointSet) Find(v byte) byte {
	if v == d.parent[v] {
		return v
	}

	d.parent[v] = d.Find(d.parent[v])
	return d.parent[v]
}

func (d *DisjointSet) Union(u, v byte) {
	u = d.Find(u)
	v = d.Find(v)
	if u == v {
		return
	}

	// update parent and rank
	if d.rank[u] < d.rank[v] {
		u, v = v, u
	}
	d.parent[v] = u

	// update rank
	if d.rank[u] == d.rank[v] {
		d.rank[u]++
	}
}
