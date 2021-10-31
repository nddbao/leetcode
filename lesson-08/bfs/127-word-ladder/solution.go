package solution

/*
	leetcode: https://leetcode.com/problems/word-ladder/
*/

/*
	TODO: improve by building adjacent based on pattern.
	For example: hot -> *ot, h*t, ho*
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var endIdx int
	if idx := findInWordList(wordList, endWord); idx >= 0 {
		endIdx = idx
	} else {
		return 0
	}

	var startIdx int
	if idx := findInWordList(wordList, beginWord); idx >= 0 {
		startIdx = idx
	} else {
		wordList = append(wordList, beginWord)
		startIdx = len(wordList) - 1
	}

	adjacents := buildAdjacents(wordList)
	if len(adjacents[startIdx]) == 0 || len(adjacents[endIdx]) == 0 {
		return 0
	}

	return bfs(adjacents, wordList, startIdx, endIdx)
}

func bfs(adjacents [][]int, wordList []string, startIdx int, endIdx int) int {
	q := &Queue{}
	q.Enque(startIdx)
	markVisited(wordList, startIdx)

	level := 1
	for q.Size() > 0 {
		size := q.Size()
		level++
		// travel all node at the same level
		for k := 0; k < size; k++ {
			node := q.Deque()
			for _, v := range adjacents[node] {
				if wordList[v] != "" {
					if v == endIdx {
						return level
					}
					q.Enque(v)
					markVisited(wordList, v)
				}
			} // end loop
		} // end loop
	} // end loop
	return 0
}

func findInWordList(wordList []string, s string) int {
	for i, v := range wordList {
		if v == s {
			return i
		}
	}
	return -1
}

func buildAdjacents(wordList []string) [][]int {
	n := len(wordList)
	a := make([][]int, len(wordList))
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if IsTransformSequence(wordList[i], wordList[j]) {
				a[i] = append(a[i], j)
				a[j] = append(a[j], i)
			}
		}
	}
	return a
}

func IsTransformSequence(s1, s2 string) bool {
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
	}

	return count == 1
}

func markVisited(wordList []string, idx int) {
	wordList[idx] = ""
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
