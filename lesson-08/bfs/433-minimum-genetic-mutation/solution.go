package solution

/*
	leetcode: https://leetcode.com/problems/minimum-genetic-mutation/
*/

/*
	We can convert this problem into graph whether each node is a gene string.
	For two nodes, they will connect to each other if they are different from each other
	by one single character.
	We build connected array and then use bfs to travel graph level by level.
	It will begin at START until reaching END.  We just return level.
	If we cannot reach END. We know that there is no way path from START to END. Just return -1.

	Time complexity:
		bfs: O(n)
		pushToBank: O(n)
		buildConnected: O(8* n^2) -> O(n^2)
	==>  O(n^2)
	Space complexity: O(n + n^2) => O(n^2)

*/

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	} else if len(bank) == 0 {
		return -1
	}

	startIdx, _ := pushToBank(&bank, start)
	endIdx, isPushed := pushToBank(&bank, end)
	if isPushed {
		return -1
	}

	isConnected := buildConnected(bank)
	visited := make([]bool, len(isConnected))

	return bfs(isConnected, visited, startIdx, endIdx)
}

func bfs(isConnected [][]int, visited []bool, startIdx, endIdx int) int {
	var level int
	visited[startIdx] = true

	q := &Queue{}
	q.Enque(startIdx)
	for q.Size() > 0 {
		size := q.Size()
		level++
		for k := 0; k < size; k++ {
			node := q.Deque()
			for _, v := range isConnected[node] {
				if visited[v] {
					continue
				}

				visited[v] = true
				if v == endIdx {
					return level
				}
				q.Enque(v)
			} // end loop
		} // end loop
	} // end loop

	return -1
}

func buildConnected(bank []string) [][]int {
	n := len(bank)
	isConnected := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !isMutation(bank[i], bank[j]) {
				continue
			}
			isConnected[i] = append(isConnected[i], j)
			isConnected[j] = append(isConnected[j], i)
		}
	}
	return isConnected
}

func pushToBank(bank *[]string, s string) (int, bool) {
	for i, v := range *bank {
		if v == s {
			return i, false
		}
	}

	*bank = append(*bank, s)
	return len(*bank) - 1, true
}

func isMutation(s1, s2 string) bool {
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}

		count++
		if count >= 2 {
			return false
		}
	}

	return count == 1
}

type Queue struct {
	a []int
}

func (q *Queue) Enque(node int) {
	q.a = append(q.a, node)
}

func (q *Queue) Deque() int {
	node := q.a[0]
	q.a = q.a[1:]
	return node
}

func (q *Queue) Size() int {
	return len(q.a)
}
