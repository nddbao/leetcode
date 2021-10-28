package solution

/*
	leetcode: https://leetcode.com/problems/binary-tree-right-side-view/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		result = append(result, queue[size-1].Val)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
