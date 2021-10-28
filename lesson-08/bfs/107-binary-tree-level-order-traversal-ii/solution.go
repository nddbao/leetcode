package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := bfs(root)
	reverseArray(result)

	return result
}

func bfs(root *TreeNode) [][]int {
	var result [][]int
	q := &Queue{}
	q.Enqueue(root)

	for q.Size() > 0 {
		var values []int
		size := q.Size()
		for i := 0; i < size; i++ {
			node := q.Dequeue()
			values = append(values, node.Val)
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		result = append(result, values)
	}
	return result
}

func reverseArray(a [][]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

type Queue struct {
	a []*TreeNode
}

func (q *Queue) Enqueue(node *TreeNode) {
	q.a = append(q.a, node)
}

func (q *Queue) Dequeue() *TreeNode {
	node := q.a[0]
	q.a = q.a[1:]
	return node
}

func (q *Queue) Size() int {
	return len(q.a)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
