package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	return dfs(root, false)
}

func dfs(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	if isLeft && root.Right == nil && root.Left == nil {
		return root.Val
	}

	return dfs(root.Right, false) + dfs(root.Left, true)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
