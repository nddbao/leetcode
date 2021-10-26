package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	var max int
	dfs(root, leftSide, &max, 0)
	dfs(root, rightSide, &max, 0)
	return max
}

const (
	leftSide  = 1
	rightSide = 2
)

func dfs(root *TreeNode, side int, max *int, count int) {
	if root == nil {
		return
	}

	if side == leftSide {
		dfs(root.Left, leftSide, max, 1)
		dfs(root.Right, rightSide, max, count+1)
	} else {
		dfs(root.Left, leftSide, max, count+1)
		dfs(root.Right, rightSide, max, 1)
	}

	if count > *max {
		*max = count
	}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
