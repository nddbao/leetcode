package solution

/*
	leetcode: https://leetcode.com/problems/reverse-linked-list/
*/

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre, head = head, tmp
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListRecursive(head *ListNode) *ListNode {
	return reverseRecursive(nil, head)
}

func reverseRecursive(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}

	tmp := cur.Next
	cur.Next = pre
	pre, cur = cur, tmp

	return reverseRecursive(pre, cur)
}
