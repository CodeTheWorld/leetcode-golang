package main

import (
	"fmt"
)

func main() {
	head := &ListNode{1, nil}
	//head.Next = &ListNode{2, nil}
	//head.Next.Next = &ListNode{3, nil}
	//head.Next.Next.Next = &ListNode{4, nil}
	//head.Next.Next.Next.Next = &ListNode{5, nil}
	res := removeNthFromEnd(head, 1)
	for nil != res {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
  思路：快慢指针
  时间复杂度：O(n)
  空间复杂度：O(1)
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{0, head} // 新的head，预防head节点被删掉的情况
	slowPtr, quickPtr := newHead, head
	for ; n > 1 && nil != quickPtr; n-- { // 快指针先走n步
		quickPtr = quickPtr.Next
	}
	if nil == quickPtr { // 快指针到头，返回原链表
		return newHead.Next
	}
	for nil != quickPtr.Next { // 快慢指针齐步走，直到慢指针到达尽头
		quickPtr = quickPtr.Next
		slowPtr = slowPtr.Next
	}
	slowPtr.Next = slowPtr.Next.Next
	return newHead.Next
}

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
