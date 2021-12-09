package solution

import "strings"

/*
	leetcode: https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/
*/

/*
	We build trie data structure.
	When inserting a folder to trie, if we found folder -> stop
	When finding folder in children
		+ if we found child is folder, just add it to result and process the other children.
		+ otherwise, we dive deep to child of child to find.


	Time complexity:
		Insert: O(N*len(subpath))
		FindFolder O(N*len(subPath))
	Space complexity: O(K) where K is sum of len all folders
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
