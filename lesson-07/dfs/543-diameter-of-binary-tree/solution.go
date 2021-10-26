package solution

/*
	leetcode: https://leetcode.com/problems/diameter-of-binary-tree/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	dfs(root, &diameter)

	return diameter
}

func dfs(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, diameter)
	right := dfs(root.Right, diameter)

	if left+right > *diameter {
		*diameter = left + right
	}

	if left > right {
		return left + 1
	}

	return right + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
