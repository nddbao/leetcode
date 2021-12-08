package solution

import "sort"

/*
	leetcode: https://leetcode.com/problems/design-search-autocomplete-system/
*/

type AutocompleteSystem struct {
	Root   *Trie
	Cursor *Trie
	Buff   []byte
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	root := &Trie{Children: make([]*Trie, 27)}
	for i := 0; i < len(sentences); i++ {
		root.Insert(sentences[i], times[i])
	}

	return AutocompleteSystem{Root: root}
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		freq := 1
		if this.Cursor != nil && this.Cursor.IsWord {
			freq += this.Cursor.Freq
		}

		word := string(this.Buff)
		this.Root.Insert(word, freq)
		this.Cursor = nil
		this.Buff = nil
		return nil
	}

	if len(this.Buff) == 0 {
		this.Cursor = this.Root.Search(c)
	} else if this.Cursor != nil {
		this.Cursor = this.Cursor.Search(c)
	}
	this.Buff = append(this.Buff, c)

	if this.Cursor == nil {
		return nil
	}

	return this.Cursor.Hot.ListString()
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */

type Trie struct {
	Children []*Trie
	Freq     int
	IsWord   bool
	Hot      HotSentences
}

func (t *Trie) Insert(word string, freq int) {
	curr := t
	sentence := Sentence{Val: word, Freq: freq}
	for i := 0; i < len(word); i++ {
		idx := getIndex(word[i])
		if curr.Children[idx] == nil {
			curr.Children[idx] = &Trie{Children: make([]*Trie, 27)}
		}
		curr.Hot = UpdateHotSentences(curr.Hot, sentence)
		curr = curr.Children[idx]
	}

	curr.Hot = UpdateHotSentences(curr.Hot, sentence)
	curr.IsWord = true
	curr.Freq = freq
}

func (t *Trie) Search(c byte) *Trie {
	idx := getIndex(c)
	return t.Children[idx]
}

func getIndex(c byte) int {
	if c == '#' {
		return -1
	}

	if c == ' ' {
		return 26
	}

	return int(c - 'a')
}

func UpdateHotSentences(h HotSentences, s Sentence) HotSentences {
	var isDuplicate bool
	for i := 0; i < len(h); i++ {
		if h[i].Val == s.Val {
			h[i] = s
			isDuplicate = true
			break
		}
	}

	if !isDuplicate {
		h = append(h, s)
	}

	sort.Sort(h)
	if len(h) < 3 {
		return h
	}

	return h[:3]
}

type Sentence struct {
	Val  string
	Freq int
}

type HotSentences []Sentence

func (h HotSentences) Len() int {
	return len(h)
}

func (h HotSentences) Less(i, j int) bool {
	if h[i].Freq == h[j].Freq {
		// sorted by asccii code (smaller to larger)
		return h[i].Val < h[j].Val
	}

	// sorted by hot degree (larger to smaller)
	return h[i].Freq > h[j].Freq
}

func (h HotSentences) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HotSentences) ListString() []string {
	var result []string
	for _, v := range h {
		result = append(result, v.Val)
	}
	return result
}
