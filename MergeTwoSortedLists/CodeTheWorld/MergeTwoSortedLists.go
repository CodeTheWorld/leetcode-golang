package CodeTheWorld

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	currentNode := head
	for nil != l1 && nil != l2 {
		if l1.Val < l2.Val {
			currentNode.Next = l1
			l1 = l1.Next
		} else {
			currentNode.Next = l2
			l2 = l2.Next
		}
		currentNode = currentNode.Next
	}
	for nil != l1 {
		currentNode.Next = l1
		currentNode = currentNode.Next
		l1 = l1.Next
	}
	for nil != l2 {
		currentNode.Next = l2
		currentNode = currentNode.Next
		l2 = l2.Next
	}
	return head.Next
}
