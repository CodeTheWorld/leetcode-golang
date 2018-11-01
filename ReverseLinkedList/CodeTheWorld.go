package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
翻转链表
时间复杂度O(n)
空间复杂度O(1)
*/
func reverseList(head *ListNode) *ListNode {
	iterator := head
	var prePtr, nextPtr *ListNode
	for nil != iterator {
		nextPtr = iterator.Next
		iterator.Next = prePtr
		prePtr = iterator
		iterator = nextPtr
	}
	return prePtr
}

func main() {
	head := &ListNode{1, nil}
	tmp := head
	for i := 2; i < 6; i++ {
		tmp.Next = &ListNode{i, nil}
		tmp = tmp.Next
	}
	res := reverseList(head)
	for nil != res {
		fmt.Println(res.Val)
		res = res.Next
	}
}
