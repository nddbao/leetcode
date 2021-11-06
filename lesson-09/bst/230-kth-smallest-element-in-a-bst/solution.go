package solution

/*
	leetcode: https://leetcode.com/problems/kth-smallest-element-in-a-bst/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	return inorderTravel(root, &k)
}

func inorderTravel(root *TreeNode, count *int) int {
	if root == nil {
		return -1
	}

	if result := inorderTravel(root.Left, count); result >= 0 {
		return result
	}

	*count--
	if *count == 0 {
		return root.Val
	}

	return inorderTravel(root.Right, count)
}
