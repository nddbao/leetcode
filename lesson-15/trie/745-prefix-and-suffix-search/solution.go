package solution

/*
	leetcode: https://leetcode.com/problems/prefix-and-suffix-search/
*/

/*
	We build trie from list words.
	For each word, we will insert all of its variant such as example below:
		"abc" -> "#abc", "a#abc", "ab#abc", "abc#abc"
	At each trie node, we will update max index.

	When calling F function, we just search word with format: suffix + "#" + prefix

	Time complexity:
		Constructor: O(len(words) * (len(maxWord)^2))
		F: O(len(prefix) + len(suffix))
	Space complexity:
		O(len(words) * (len(maxWord)^2))
*/

type WordFilter struct {
	words []string
	root  *Trie
}

func Constructor(words []string) WordFilter {
	root := NewTrie()
	for i := 0; i < len(words); i++ {
		for j := 0; j <= len(words[i]); j++ {
			w := words[i][j:] + "#" + words[i]
			root.Insert(w, i)
		} // end loop
	} // end loop

	return WordFilter{words: words, root: root}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	search := suffix + "#" + prefix
	cur := this.root.Find(search)
	if cur != nil {
		return cur.MaxIdx
	}

	return -1
}

func NewTrie() *Trie {
	return &Trie{Children: make([]*Trie, 27), MaxIdx: -1}
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */

type Trie struct {
	Children []*Trie
	IsWord   bool
	MaxIdx   int
}

func (t *Trie) Insert(word string, wordIdx int) {
	cur := t
	for i := 0; i < len(word); i++ {
		idx := getIndex(word[i])
		if cur.Children[idx] == nil {
			cur.Children[idx] = NewTrie()
		}

		cur.Children[idx].SetMaxIdx(wordIdx)
		cur = cur.Children[idx]
	}

	cur.IsWord = true
}

func (t *Trie) Find(prefix string) *Trie {
	cur := t
	for i := 0; i < len(prefix); i++ {
		idx := getIndex(prefix[i])
		if cur.Children[idx] == nil {
			return nil
		}
		cur = cur.Children[idx]
	}

	return cur
}

func (t *Trie) SetMaxIdx(wordIdx int) {
	if t.MaxIdx < wordIdx {
		t.MaxIdx = wordIdx
	}
}

func getIndex(c byte) int {
	if c == '#' {
		return 26
	}

	return int(c - 'a')
}
