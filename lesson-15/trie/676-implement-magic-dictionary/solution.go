package solution

/*
	leetcode: https://leetcode.com/problems/implement-magic-dictionary/
*/

type MagicDictionary struct {
	root *Node
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		root: &Node{Children: make([]*Node, 26)},
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		this.root.InsertWord(word)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for i := 0; i < len(searchWord); i++ {
		if this.root.Search(searchWord, i, 0) {
			return true
		}
	} // end loop

	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

type Node struct {
	Children []*Node
	End      bool
}

func (n *Node) Search(s string, skipIdx, k int) bool {
	if k == len(s) {
		return n.End
	}

	if skipIdx == k {
		for i := 0; i < len(n.Children); i++ {
			child := n.Children[i]
			if child != nil && child.Search(s, skipIdx, k+1) && i != n.GetIndex(s[k]) {
				return true
			}
		} // end loop

		return false
	}

	idx := n.GetIndex(s[k])
	child := n.Children[idx]
	if child != nil {
		return child.Search(s, skipIdx, k+1)
	}

	return false
}

func (n *Node) InsertWord(word string) {
	curr := n
	for i := 0; i < len(word); i++ {
		curr = curr.Add(word[i])
	}
	curr.End = true
}

func (n *Node) Add(c byte) *Node {
	idx := n.GetIndex(c)
	if n.Children[idx] == nil {
		n.Children[idx] = &Node{Children: make([]*Node, 26)}
	}
	return n.Children[idx]
}

func (n *Node) GetIndex(c byte) int {
	return int(c - 'a')
}
