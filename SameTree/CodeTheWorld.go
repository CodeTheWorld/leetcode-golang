package main

import "fmt"

func main() {
	t1 := &TreeNode{1, nil, nil}
	t1.Left = &TreeNode{2, nil, nil}
	t1.Right = &TreeNode{1, nil, nil}
	t2 := &TreeNode{1, nil, nil}
	t2.Left = &TreeNode{2, nil, nil}
	t2.Right = &TreeNode{1, nil, nil}
	res := isSameTree1(t1, t2)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
  思路：递归
  时间复杂度：O(n)
  空间复杂度：O(1)
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if nil == p && nil == q {
		return true
	}
	if nil == p || nil == q || p.Val != q.Val {
		return false
	}
	leftRes := isSameTree(p.Left, q.Left)
	rightRes := isSameTree(p.Right, q.Right)
	return leftRes && rightRes
}

type stack struct {
	BinTree []*TreeNode
	Top     int
}

func push(nodeStack *stack, node *TreeNode) {
	nodeStack.BinTree = append(nodeStack.BinTree, node)
	nodeStack.Top++
}

func pop(nodeStack *stack) *TreeNode {
	var res *TreeNode
	nodeStack.Top--
	if nodeStack.Top < 0 {
		res = nil
	} else {
		res = nodeStack.BinTree[nodeStack.Top]
	}
	nodeStack.BinTree = nodeStack.BinTree[:nodeStack.Top]
	return res
}

/**
  思路：利用栈，迭代
  时间复杂度：O(n)
  空间复杂度：O(n)
*/
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	myStack := &stack{[]*TreeNode{}, 0}
	push(myStack, p)
	push(myStack, q)

	for 0 < myStack.Top {
		left := pop(myStack)
		right := pop(myStack)
		if nil == left && nil == right {
			continue
		}
		if nil == left || nil == right || left.Val != right.Val {
			return false
		}
		push(myStack, left.Left)
		push(myStack, right.Left)
		push(myStack, left.Right)
		push(myStack, right.Right)
	}
	return true
}
