package solution

import "strings"

/*
	leetcode: https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/
*/

func removeSubfolders(folder []string) []string {
	root := &Trie{Children: make(map[string]*Trie)}
	for _, v := range folder {
		root.Insert(v)
	}

	return root.FindFolder("")
}

type Trie struct {
	Children map[string]*Trie
	IsFolder bool
}

func (t *Trie) Insert(folder string) {
	cur := t
	list := strings.Split(folder, "/")
	for i := 1; i < len(list); i++ {
		k := list[i]
		if _, ok := cur.Children[k]; !ok {
			cur.Children[k] = &Trie{Children: make(map[string]*Trie)}
		}

		cur = cur.Children[k]
		if cur.IsFolder {
			return
		}
	}

	cur.IsFolder = true
}

func (t *Trie) FindFolder(prefix string) []string {
	var result []string
	for key, child := range t.Children {
		s := prefix + "/" + key
		if child.IsFolder {
			result = append(result, s)
			continue
		}

		subFolders := child.FindFolder(s)
		if len(subFolders) > 0 {
			result = append(result, subFolders...)
		}
	}

	return result
}
