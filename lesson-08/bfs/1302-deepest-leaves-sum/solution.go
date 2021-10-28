package solution

/*
	leetcode: https://leetcode.com/problems/deepest-leaves-sum/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	var sum int
	for len(queue) > 0 {
		size := len(queue)
		sum = 0

		for i := 0; i < size; i++ {
			node := queue[i]
			queue = queue[1:]

			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return sum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
