package solution

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	var count int
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			dfs(isConnected, visited, i)
			count++
		}
	}

	return count
}

func dfs(isConnected [][]int, visited []bool, city int) {
	visited[city] = true
	for i := 0; i < len(isConnected); i++ {
		if visited[i] || isConnected[city][i] == 0 {
			continue
		}
		dfs(isConnected, visited, i)
	}
}
