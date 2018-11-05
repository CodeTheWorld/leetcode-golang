package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{2, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{3, nil}
	l2 := &ListNode{5, nil}
	l2.Next = &ListNode{6, nil}
	l2.Next = &ListNode{4, nil}
	res := addTwoNumbers(l1, l2)
	for nil != res {
		fmt.Println(res.Val)
		res = res.Next
	}
}

/**
  思路：大数相加，遍历，注意进位
  时间复杂度：O(n)
  空间复杂度：O(n)
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	bit, flag, iterator := 0, 0, sum // bit 每位和，flag 进位标志 iterator 迭代器
	for nil != l1 || nil != l2 {
		bit = flag
		if nil != l1 {
			bit += l1.Val
			l1 = l1.Next
		}
		if nil != l2 {
			bit += l2.Val
			l2 = l2.Next
		}
		flag = bit / 10
		iterator.Next = &ListNode{bit % 10, nil}
		iterator = iterator.Next
	}
	return sum.Next
}
