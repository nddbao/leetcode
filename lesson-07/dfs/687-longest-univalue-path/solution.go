package solution

/*
	leetcode: https://leetcode.com/problems/longest-univalue-path/
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	var max int
	dfs(root, &max)
	return max
}

func dfs(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, max)
	right := dfs(root.Right, max)

	if root.Left != nil && root.Left.Val == root.Val &&
		root.Right != nil && root.Right.Val == root.Val {

		if *max < left+right+2 {
			*max = left + right + 2
		}

		if left > right {
			return left + 1
		}
		return right + 1
	}

	if root.Left != nil && root.Left.Val == root.Val {
		if *max < left+1 {
			*max = left + 1
		}

		return left + 1
	}

	if root.Right != nil && root.Right.Val == root.Val {
		if *max < right+1 {
			*max = right + 1
		}

		return right + 1
	}

	return 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
