package solution

/*
	leetcode: https://leetcode.com/problems/sum-of-left-leaves/
*/

/*
	We will use dfs to travel all of node in the binary tree.
	We also have a parameter to know current is left or right.
	If current node is left side, we will add it to our sum.

	Time complexity: O(n) where n is number node of the tree
	Space complexity: O(h) where h i maximum height of the tree
*/

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
