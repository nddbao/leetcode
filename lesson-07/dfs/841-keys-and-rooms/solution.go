package solution

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	dfs(rooms, visited, 0)

	for _, v := range visited {
		if v == false {
			return false
		}
	}

	return true
}

func dfs(rooms [][]int, visited []bool, roomIdx int) {
	visited[roomIdx] = true
	for _, v := range rooms[roomIdx] {
		if visited[v] {
			continue
		}

		dfs(rooms, visited, v)
	}
}
