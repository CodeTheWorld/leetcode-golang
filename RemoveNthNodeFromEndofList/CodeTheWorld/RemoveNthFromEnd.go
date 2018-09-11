package CodeTheWorld

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if 0 >= n || nil == head {
		return head
	}

	current, first := &ListNode{0, head}, head
	tmpHead := current

	for nil != first.Next && n > 1 {
		first = first.Next
		n--
	}
	for nil != first.Next {
		first = first.Next
		current = current.Next
	}
	if 1 < n {
		return head
	}
	current.Next = current.Next.Next
	return tmpHead.Next
}
