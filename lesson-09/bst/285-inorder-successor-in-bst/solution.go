package solution

/*
	leetcode: https://leetcode.com/problems/inorder-successor-in-bst/
*/

/*
	We just search from root to find. Comparing current node value to p value
	There are there 3  cases:
		+ node.Val <= p.Val: we know that we need to find bigger value by moving to right side.
		+ node.Val > p.Val : we know that current node can be our result.
							Maybe there is a smaller value on the left side that still is greater than p.Val
							Let try to move to left side to check it.

	Time complexity: O(N)
	Space Complexity: O(1)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var result *TreeNode
	for root != nil {
		if root.Val > p.Val {
			result = root
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return result
}
