package main

import "fmt"

func main() {
	//head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	head := &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}}
	res := insertionSortList(head)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
思路：插入排序
时间复杂度：O(n2)
空间复杂度：O(1)
*/
func insertionSortList(head *ListNode) *ListNode {
	newHead := &ListNode{0, nil}
	var node *ListNode
	for nil != head {
		node = head
		head = head.Next
		for iterator := newHead; iterator != nil; iterator = iterator.Next {
			if nil == iterator.Next || iterator.Next.Val >= node.Val {
				node.Next = iterator.Next
				iterator.Next = node
				break
			}
		}
	}
	return newHead.Next
}
