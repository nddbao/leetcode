package solution

import "math"

/*
	leetcode: https://leetcode.com/problems/minimum-depth-of-binary-tree/
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return math.MaxInt32
	}

	if root.Right == nil && root.Left == nil {
		return 1
	}

	left := dfs(root.Left)
	right := dfs(root.Right)
	if left > right {
		return right + 1
	}

	return left + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
