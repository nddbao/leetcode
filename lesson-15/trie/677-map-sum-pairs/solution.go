package solution

/*
	leetcode: https://leetcode.com/problems/map-sum-pairs/
*/

type MapSum struct {
	trie  *Trie
	cache map[string]int
}

func Constructor() MapSum {
	trie := &Trie{Children: make([]*Trie, 26)}
	cache := make(map[string]int)
	return MapSum{trie: trie, cache: cache}
}

func (this *MapSum) Insert(key string, val int) {
	diff := val
	v, ok := this.cache[key]
	if ok {
		diff -= v
	}

	this.trie.Add(key, diff)
	this.cache[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	return this.trie.SumPrefix(prefix)
}

func getIndex(c byte) int {
	return int(c - 'a')
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

type Trie struct {
	Children []*Trie
	Sum      int
}

func (t *Trie) Add(key string, val int) {
	curr := t
	curr.Sum += val
	for i := 0; i < len(key); i++ {
		idx := getIndex(key[i])
		if curr.Children[idx] == nil {
			curr.Children[idx] = &Trie{Children: make([]*Trie, 26)}
		}
		curr = curr.Children[idx]
		curr.Sum += val
	}
}

func (t *Trie) SumPrefix(prefix string) int {
	curr := t
	for i := 0; i < len(prefix); i++ {
		idx := getIndex(prefix[i])
		if curr.Children[idx] == nil {
			return 0
		}
		curr = curr.Children[idx]
	}
	return curr.Sum
}
