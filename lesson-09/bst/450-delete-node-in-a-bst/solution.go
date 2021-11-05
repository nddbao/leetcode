package solution

/*
	leetcode: https://leetcode.com/problems/delete-node-in-a-bst/
*/

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	dumb := &TreeNode{Val: math.MaxInt32}
	dumb.Left = root
	parent, side := parentOfTarget(dumb, dumb.Left, key, SideLeft)
	if parent == nil {
		return dumb.Left
	}

	node := getNode(parent, side)
	for !isLeaf(node) {
		if node.Right == nil || node.Left == nil {
			child := node.Right
			if node.Left != nil {
				child = node.Left
			}
			setChild(parent, child, side)
			return dumb.Left
		}

		parent, side = parentOfMin(node, node.Right, SideRight)
		tmp := getNode(parent, side)
		node.Val, tmp.Val = tmp.Val, node.Val
		node = tmp
	}

	deleteLeaf(parent, side)

	return dumb.Left
}

func setChild(parent, node *TreeNode, side int) {
	if side == SideLeft {
		parent.Left = node
	} else {
		parent.Right = node
	}
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func parentOfMin(parent, node *TreeNode, side int) (*TreeNode, int) {
	for node.Left != nil {
		parent, node = node, node.Left
		side = SideLeft
	}
	return parent, side
}

func deleteLeaf(node *TreeNode, side int) {
	if side == SideLeft {
		node.Left = nil
	} else {
		node.Right = nil
	}
}

const (
	SideLeft  int = 1
	SideRight int = 2
)

func parentOfTarget(parent, node *TreeNode, key int, side int) (*TreeNode, int) {
	for node != nil {
		if node.Val == key {
			return parent, side
		}

		if node.Val < key {
			parent, node = node, node.Right
			side = SideRight
		} else {
			parent, node = node, node.Left
			side = SideLeft
		}
	}

	return nil, 0
}

func getNode(parent *TreeNode, side int) *TreeNode {
	if side == SideLeft {
		return parent.Left
	}
	return parent.Right
}
