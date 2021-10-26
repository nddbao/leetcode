package solution

/*
	leetcode: https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
*/

func countComponents(n int, edges [][]int) int {
	mEdges := make(map[int][]int)
	for _, v := range edges {
		v1 := v[0]
		v2 := v[1]
		mEdges[v1] = append(mEdges[v1], v2)
		mEdges[v2] = append(mEdges[v2], v1)
	}

	visited := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(visited, mEdges, i)
			count++
		}
	}

	return count
}

func dfs(visited []bool, mEdges map[int][]int, node int) {
	if visited[node] {
		return
	}

	visited[node] = true

	connectedNodes, ok := mEdges[node]
	if !ok {
		return
	}

	for _, v := range connectedNodes {
		dfs(visited, mEdges, v)
	}
}
