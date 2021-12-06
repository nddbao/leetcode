package solution

/*
	leetcode: https://leetcode.com/problems/implement-trie-prefix-tree/
*/

type Trie struct {
	Children []*Trie
	IsWord   bool
}

func Constructor() Trie {
	return Trie{Children: make([]*Trie, 26)}
}

func (this *Trie) Insert(word string) {
	curr := this
	for i := 0; i < len(word); i++ {
		idx := getIndex(word[i])
		if curr.Children[idx] == nil {
			curr.Children[idx] = &Trie{Children: make([]*Trie, 26)}
		}
		curr = curr.Children[idx]
	}
	curr.IsWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for i := 0; i < len(word); i++ {
		idx := getIndex(word[i])
		if curr.Children[idx] == nil {
			return false
		}
		curr = curr.Children[idx]
	}

	return curr.IsWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for i := 0; i < len(prefix); i++ {
		idx := getIndex(prefix[i])
		if curr.Children[idx] == nil {
			return false
		}
		curr = curr.Children[idx]
	}
	return true
}

func getIndex(c byte) int {
	return int(c - 'a')
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
