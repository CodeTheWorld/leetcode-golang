package MergeKSortedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
暴力解法，依次比较大小，将小的插入到结果链表中
时间复杂度:O(k*n)
*/
func MergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	var minNode *ListNode
	res := &ListNode{0, nil}
	tmp := res

	for {
		minNode = nil
		minIndex := 0
		for i := 0; i < length; i++ {
			if nil == lists[i] {
				continue
			}
			if nil == minNode {
				minNode = lists[i]
				minIndex = i
			}
			if minNode.Val > lists[i].Val {
				minNode = lists[i]
				minIndex = i
			}
		}
		if nil == minNode {
			break
		}
		tmp.Next = lists[minIndex]
		tmp = tmp.Next
		lists[minIndex] = lists[minIndex].Next
	}

	return res.Next
}

/**
分治法
时间复杂度:O(log2(k)*2*n)
*/
func MergeKLists1(lists []*ListNode) *ListNode {
	length := len(lists)
	n := length
	if 0 == n {
		return nil
	}
	for n > 1 {
		for left, right := 0, n-1; left < right; {
			lists[left] = mergeTwoLists(lists[left], lists[right])
			left++
			right--

			if left == right {
				n = n/2 + 1
			} else if left > right {
				n = n / 2
			}
		}
	}

	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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
