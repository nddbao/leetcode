package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func robOptimize(root *TreeNode) int {
	return max(dfsOptimize(root))
}

func dfsOptimize(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0} // [rob, notRob]
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val, 0}
	}

	left := dfsOptimize(root.Left)
	right := dfsOptimize(root.Right)

	rob := root.Val + left[1] + right[1]
	notRob := max(left) + max(right)

	return []int{rob, notRob}
}

func max(a []int) int {
	if a[0] > a[1] {
		return a[0]
	}
	return a[1]
}
