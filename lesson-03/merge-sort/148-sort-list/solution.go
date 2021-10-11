package solution

/*
	leetcode: https://leetcode.com/problems/sort-list/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	preMid := preMidNode(head)
	l1, l2 := head, preMid.Next
	preMid.Next = nil

	l1 = sortList(l1)
	l2 = sortList(l2)

	return mergeSortedLinkedList(l1, l2)
}

func preMidNode(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func mergeSortedLinkedList(l1, l2 *ListNode) *ListNode {
	if l2 == nil {
		return l1
	}

	dummyHead := &ListNode{}
	curr := dummyHead
	for l1 != nil && l2 != nil {
		if l2.Val < l1.Val {
			l1, l2 = l2, l1
		}
		curr.Next = l1
		curr, l1 = curr.Next, l1.Next
	}

	if l2 != nil {
		curr.Next = l2
	}

	return dummyHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
