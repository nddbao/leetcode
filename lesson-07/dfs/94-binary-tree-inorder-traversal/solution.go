package solution

/*
	leetcode: https://leetcode.com/problems/binary-tree-inorder-traversal/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var a []int
	helper(root, &a)
	return a
}

func helper(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}

	helper(root.Left, a)
	*a = append(*a, root.Val)
	helper(root.Right, a)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
