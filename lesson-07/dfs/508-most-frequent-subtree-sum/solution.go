package solution

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	mSum := make(map[int]int)
	dfs(root, mSum)

	max := FindMax(mSum)

	var result []int
	for k, v := range mSum {
		if v == max {
			result = append(result, k)
		}
	}
	return result
}

func FindMax(mSum map[int]int) int {
	var max = math.MinInt32
	for _, v := range mSum {
		if max < v {
			max = v
		}
	}
	return max
}

func dfs(root *TreeNode, mSum map[int]int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, mSum)
	right := dfs(root.Right, mSum)
	s := left + right + root.Val
	mSum[s]++

	return s
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
