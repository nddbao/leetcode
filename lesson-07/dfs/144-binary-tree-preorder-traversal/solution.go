package solution

/*
	leetcode: https://leetcode.com/problems/binary-tree-preorder-traversal/
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var a []int
	helper(root, &a)
	return a
}

func helper(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}

	*a = append(*a, root.Val)
	helper(root.Left, a)
	helper(root.Right, a)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
