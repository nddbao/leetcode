package solution

/*
	leetcode: https://leetcode.com/problems/linked-list-cycle-ii/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	// check cycle
	var hasCycle bool
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil
	}

	entry := head
	for entry != slow {
		entry, slow = entry.Next, slow.Next
	}

	return entry
}

type ListNode struct {
	Val  int
	Next *ListNode
}
