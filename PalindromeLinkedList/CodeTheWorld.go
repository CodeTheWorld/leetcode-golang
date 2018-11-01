package PalindromeLinkedList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test() {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{2, nil}
	l1.Next.Next.Next = &ListNode{1, nil}

	res := isPalindrome(l1)
	fmt.Println(res)
}

func isPalindrome(head *ListNode) bool {
	quickPtr, slowPtr := head, head
	var prePtr, nextPtr *ListNode // 前一个slow指针

	for nil != quickPtr && nil != quickPtr.Next {
		quickPtr = quickPtr.Next.Next
		nextPtr = slowPtr.Next
		slowPtr.Next = prePtr
		prePtr = slowPtr
		slowPtr = nextPtr
	}
	if nil == quickPtr { // even
		nextPtr = slowPtr
	} else { // odd
		nextPtr = slowPtr.Next
	}
	for nil != nextPtr {
		if nextPtr.Val != prePtr.Val {
			return false
		}
		nextPtr = nextPtr.Next
		prePtr = prePtr.Next
	}

	return true
}
