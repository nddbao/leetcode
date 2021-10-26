package solution

/*
	leetcode: https://leetcode.com/problems/maximum-depth-of-binary-tree/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)
	if right > left {
		return right + 1
	}
	return left + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
