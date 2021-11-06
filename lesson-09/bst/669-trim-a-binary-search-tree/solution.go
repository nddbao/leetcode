package solution

/*
	leetcode: https://leetcode.com/problems/trim-a-binary-search-tree/
*/

/*
	We check value at each node comparing with range [low..high]
	There are three 3 cases:
	+ node.Val < low : all values of left side and node is not satisfied.
		=> keep only right side, move to right side to trim.
	+ node.Val > high: all values of right side and node is not satisfied.
		=> keep only left side, move to left side to trim.
    + node.Val in [low..high]: move to down to left and right to trim.

	Time complexity: O(n)
	Space complexity: O(n)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
