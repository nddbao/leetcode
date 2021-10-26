package solution

/*
	leetcode: https://leetcode.com/problems/house-robber-iii/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	cache := make(map[*TreeNode][]int)
	a := dfs(root, cache, true)
	b := dfs(root, cache, false)
	return findMax(a, b)
}

func dfs(root *TreeNode, cache map[*TreeNode][]int, isRob bool) int {
	if root == nil {
		return 0
	}

	v, ok := checkCache(cache, root, isRob)
	if ok {
		return v
	}

	var money int
	if isRob {
		money = dfs(root.Left, cache, false) + dfs(root.Right, cache, false) + root.Val
	} else {
		m1 := dfs(root.Left, cache, true) + dfs(root.Right, cache, false)
		m2 := dfs(root.Left, cache, false) + dfs(root.Right, cache, true)
		m3 := dfs(root.Left, cache, true) + dfs(root.Right, cache, true)
		m4 := dfs(root.Left, cache, false) + dfs(root.Right, cache, false)
		money = findMax(m1, m2, m3, m4)
	}

	addToCache(cache, root, isRob, money)

	return money
}

func checkCache(cache map[*TreeNode][]int, node *TreeNode, isRob bool) (int, bool) {
	v, ok := cache[node]
	if !ok {
		return 0, false
	}

	if isRob && v[0] >= 0 {
		return v[0], true
	}

	if !isRob && v[1] >= 0 {
		return v[1], true
	}

	return 0, false
}

func addToCache(cache map[*TreeNode][]int, node *TreeNode, isRob bool, money int) {
	v, ok := cache[node]
	if !ok {
		v = []int{-1, -1}
	}

	if isRob {
		v[0] = money
	} else {
		v[1] = money
	}

	cache[node] = v
}

func findMax(a ...int) int {
	max := a[0]
	for _, v := range a {
		if max < v {
			max = v
		}
	}

	return max
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
