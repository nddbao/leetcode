package solution

/*
	leetcode: https://leetcode.com/problems/longest-word-in-dictionary/
*/

func longestWord(words []string) string {
	root := &Trie{Children: make([]*Trie, 26)}
	for _, word := range words {
		root.Insert(word)
	}

	result := root.FindLongest()
	return string(result)
}

type Trie struct {
	Children []*Trie
	IsWord   bool
}

func (t *Trie) Insert(word string) {
	curr := t
	for i := 0; i < len(word); i++ {
		idx := getIndex(word[i])
		if curr.Children[idx] == nil {
			curr.Children[idx] = &Trie{Children: make([]*Trie, 26)}
		}
		curr = curr.Children[idx]
	}
	curr.IsWord = true
}

func (t *Trie) FindLongest() []byte {
	var max []byte
	var maxIdx = -1
	for i := 0; i < len(t.Children); i++ {
		if t.Children[i] == nil || t.Children[i].IsWord == false {
			continue
		}

		val := t.Children[i].FindLongest()
		if maxIdx == -1 || len(val) > len(max) {
			max = val
			maxIdx = i
		}
	}

	if maxIdx == -1 {
		return nil
	}

	max = append([]byte{byte(maxIdx) + 'a'}, max...)
	return max
}

func getIndex(c byte) int {
	return int(c - 'a')
}
