package solution

/*
	leetcode: https://leetcode.com/problems/number-of-islands
*/

/*
	We apply union-find with path compression.
	Each cell "1" will be consider as a subset.
	We try to merge subset if they are next to each other.
	Finally, we count how many subsets in our disjoint-set.

	Time complexity: O(m*n)
		Apply  union with path compression so Find, Union, MakeSet ~ O(1)
	Space complexity: O(m*n)

*/
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	visited := NewVisited(grid)
	ds := NewDisjointSet(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isIsland(grid, i, j) && !visited[i][j] {
				updateIslands(grid, visited, ds, i, j)
			}
		} /// end loop
	} // end loop

	return ds.NumberOfSet()
}

func updateIslands(grid [][]byte, visited [][]bool, ds *DisjointSet, i, j int) {
	n := len(grid[0])
	u := getIndex(n, i, j)
	ds.MakeSet(u)

	visited[i][j] = true
	for _, d := range directions {
		dx, dy := i+d[0], j+d[1]
		if !canUnion(grid, visited, dx, dy) {
			continue
		}

		v := getIndex(n, dx, dy)
		v = ds.MakeSet(v)
		ds.Union(u, v)
	} // end loop
}

func getIndex(n int, i, j int) int {
	return i*n + j
}

var directions = [][]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

func NewVisited(grid [][]byte) [][]bool {
	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	return visited
}

func isValidIdx(grid [][]byte, i, j int) bool {
	m := len(grid)
	n := len(grid[0])
	return i >= 0 && i < m && j >= 0 && j < n
}

func isIsland(grid [][]byte, i, j int) bool {
	return grid[i][j] == '1'
}

func canUnion(grid [][]byte, visited [][]bool, i, j int) bool {
	return isValidIdx(grid, i, j) && isIsland(grid, i, j) && !visited[i][j]
}

func NewDisjointSet(n, m int) *DisjointSet {
	parent := make([]int, n*m)
	rank := make([]int, n*m)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}

	return &DisjointSet{parent: parent, rank: rank}
}

type DisjointSet struct {
	parent      []int
	rank        []int
	numberOfSet int
}

func (d *DisjointSet) MakeSet(v int) int {
	if d.parent[v] >= 0 {
		return d.Find(v)
	}

	d.parent[v] = v

	// update number of set
	d.numberOfSet++
	return v
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

	// update parent and rank
	if d.rank[u] < d.rank[v] {
		u, v = v, u
	}
	d.parent[v] = u

	if d.rank[u] == d.rank[v] {
		d.rank[u] += 1
	}

	// update number of set
	d.numberOfSet--
}

func (d *DisjointSet) NumberOfSet() int {
	return d.numberOfSet
}

func (d *DisjointSet) Parent(u int) int {
	return d.parent[u]
}
