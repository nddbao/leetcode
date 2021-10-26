package solution

/*
	leetcode: https://leetcode.com/problems/binary-tree-postorder-traversal/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var a []int
	helper(root, &a)
	return a
}

func helper(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}

	helper(root.Left, a)
	helper(root.Right, a)
	*a = append(*a, root.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
