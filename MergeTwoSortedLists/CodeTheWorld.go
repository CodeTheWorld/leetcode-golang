package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{4, nil}
	l2 := &ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}

	res := mergeTwoLists(l1, l2)
	for nil != res {
		fmt.Println(res.Val)
		res = res.Next
	}
}

/**
  思路：链表遍历
  时间复杂度：O(n)
  空间复杂度：O(1)
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	ptr := head

	for nil != l1 && nil != l2 { // 两个链表都不为空
		if l1.Val <= l2.Val {
			ptr.Next = l1
			l1 = l1.Next
		} else {
			ptr.Next = l2
			l2 = l2.Next
		}
		ptr = ptr.Next
	}

	for nil != l1 { // l1不为空
		ptr.Next = l1
		l1 = l1.Next
		ptr = ptr.Next
	}
	for nil != l2 { // l2不为空
		ptr.Next = l2
		l2 = l2.Next
		ptr = ptr.Next
	}

	return head.Next
}
