package solution

/*
	leetcode: https://leetcode.com/problems/inorder-successor-in-bst/
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
