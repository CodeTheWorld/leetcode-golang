package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	quickPtr, slowPtr := head, head
	for nil != quickPtr && nil != quickPtr.Next {
		slowPtr = slowPtr.Next
		quickPtr = quickPtr.Next.Next
	}
	return slowPtr
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	res := middleNode(head)
	fmt.Println(res.Val)
}
