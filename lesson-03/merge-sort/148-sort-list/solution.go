package solution

/*
	leetcode: https://leetcode.com/problems/sort-list/
*/

/*
	We use divide list into 2 sublist, sort them recursively and merge two sorted list.
	For dividing into 2 sublist: we will find middle node and split list based on this.
	For merging two 2 sorted list l1, l2: compare l1 and l2 value. Choose smaller one and
	update it to our new list. Keep moving until list 1 or list 2 don't have any node.

	Time complexity: O(nlogn) (find middle O(n) and merge O(nlogn)
	Space complexity: O(1) without counting for call stack
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
