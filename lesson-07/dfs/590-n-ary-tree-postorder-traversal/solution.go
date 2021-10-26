package solution

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	var a []int
	helper(root, &a)
	return a
}

func helper(root *Node, a *[]int) {
	if root == nil {
		return
	}

	for _, v := range root.Children {
		helper(v, a)
	}

	*a = append(*a, root.Val)
}

type Node struct {
	Val      int
	Children []*Node
}
