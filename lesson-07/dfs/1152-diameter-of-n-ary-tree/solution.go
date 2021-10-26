package solution

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func diameter(root *Node) int {
	if root == nil {
		return 0
	}
	var max int
	dfs(root, &max)
	return max
}

func dfs(root *Node, max *int) int {
	if len(root.Children) == 0 {
		return 1
	}

	max1, max2 := 0, 0
	for _, node := range root.Children {
		level := dfs(node, max)
		if max1 <= level {
			max1, max2 = level, max1
		} else if max2 < level {
			max2 = level
		}
	}

	if max1+max2 > *max {
		*max = max1 + max2
	}

	return max1 + 1
}

type Node struct {
	Val      int
	Children []*Node
}
