package solution

/*
	leetcode: https://leetcode.com/problems/n-ary-tree-level-order-traversal/
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	q := &Queue{}
	q.Enqueue(root)
	for q.Size() > 0 {
		var values []int
		size := q.Size()
		for i := 0; i < size; i++ {
			node := q.Dequeue()
			values = append(values, node.Val)
			for _, v := range node.Children {
				q.Enqueue(v)
			}
		}
		result = append(result, values)
	}

	return result
}

type Queue struct {
	a []*Node
}

func (q *Queue) Enqueue(node *Node) {
	q.a = append(q.a, node)
}

func (q *Queue) Dequeue() *Node {
	node := q.a[0]
	q.a = q.a[1:]
	return node
}

func (q *Queue) Size() int {
	return len(q.a)
}

type Node struct {
	Val      int
	Children []*Node
}
