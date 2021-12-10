package solution

/*
	leetcode: https://leetcode.com/problems/stream-of-characters/
*/

type StreamChecker struct {
	root *Trie
	buff []byte
}

func Constructor(words []string) StreamChecker {
	root := &Trie{Children: make([]*Trie, 26)}
	for _, w := range words {
		root.InsertReverse(w)
	}

	return StreamChecker{root: root}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.buff = append(this.buff, letter)
	return this.root.Find(this.buff) != nil
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */

type Trie struct {
	Children []*Trie
	IsWord   bool
}

func (t *Trie) InsertReverse(word string) {
	cur := t
	for i := len(word) - 1; i >= 0; i-- {
		idx := getIndex(word[i])
		if cur.Children[idx] == nil {
			cur.Children[idx] = &Trie{Children: make([]*Trie, 26)}
		}

		cur = cur.Children[idx]
	}
	cur.IsWord = true
}

func (t *Trie) Find(suffix []byte) *Trie {
	cur := t
	for i := len(suffix) - 1; i >= 0; i-- {
		idx := getIndex(suffix[i])
		if cur.Children[idx] == nil {
			return nil
		}

		if cur.Children[idx].IsWord {
			return cur.Children[idx]
		}

		cur = cur.Children[idx]
	}

	return nil
}

func getIndex(c byte) int {
	return int(c - 'a')
}
