package solution

import "sort"

/*
	leetcode: https://leetcode.com/problems/search-suggestions-system/
*/

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	root := &Trie{
		Children: make([]*Trie, 26),
		Start:    0,
		End:      len(products) - 1,
	}

	for i := 0; i < len(products); i++ {
		root.Insert(products[i], i)
	}

	result := make([][]string, len(searchWord))
	curr := root
	for i := 0; i < len(searchWord); i++ {
		curr = curr.Search(searchWord[i : i+1])
		if curr == nil {
			break
		}

		result[i] = getTopThree(products, curr.Start, curr.End)
	}
	return result
}

func getTopThree(products []string, i, j int) []string {
	if j-i+1 > 3 {
		j = i + 2
	}

	return products[i : j+1]
}

type Trie struct {
	Children   []*Trie
	Start, End int
}

func (t *Trie) Insert(word string, wordIdx int) {
	curr := t
	for i := 0; i < len(word); i++ {
		idx := getIndex(word[i])
		if curr.Children[idx] == nil {
			curr.Children[idx] = &Trie{
				Children: make([]*Trie, 26),
				Start:    wordIdx,
				End:      wordIdx,
			}
		} else {
			curr.Children[idx].End = wordIdx
		}
		curr = curr.Children[idx]
	}
}

func (t *Trie) Search(prefix string) *Trie {
	curr := t
	for i := 0; i < len(prefix); i++ {
		idx := getIndex(prefix[i])
		if curr.Children[idx] == nil {
			return nil
		}
		curr = curr.Children[idx]
	}

	return curr
}

func getIndex(c byte) int {
	return int(c - 'a')
}
