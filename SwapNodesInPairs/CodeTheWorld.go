package SwapNodesInPairs

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var tmp *ListNode
	pre := &ListNode{0, head}
	current := pre
	for nil != current.Next && nil != current.Next.Next {
		tmp = current.Next
		current.Next = tmp.Next
		tmp.Next = current.Next.Next
		current.Next.Next = tmp
		current = current.Next.Next
	}

	return pre.Next
}

func Test() {
	l1 := &ListNode{-1, nil}
	l1.Next = &ListNode{5, nil}
	l1.Next.Next = &ListNode{11, nil}
	res := swapPairs(l1)
	for nil != res {
		fmt.Println(res.Val)
		res = res.Next
	}
}
