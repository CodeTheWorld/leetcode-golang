package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{5, nil}
	l2 := &ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}
	l3 := &ListNode{2, nil}
	l3.Next = &ListNode{6, nil}
	lists := []*ListNode{l1, l2, l3}
	res := mergeKLists(lists)
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
  思路：分支法
  时间复杂度：O(nklogk) 最外层遍历复杂度 O(logk)；每次循环体的时间复杂度O(nk/2)(k/2次遍历2n个元素...)
  空间复杂度：O(1)
*/
func mergeKLists(lists []*ListNode) *ListNode {
	count := len(lists)
	if 0 == count {
		return nil
	}

	for count > 1 {
		for i := 0; i < count/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[count-i-1])
		}
		count = (count + 1) / 2
	}

	return lists[0]
}

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

func mergeKLists1(lists []*ListNode) *ListNode {
}
