package main

import "fmt"

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}
	res := isSymmetric1(root)
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
func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return isSymmetricRecursion(root.Left, root.Right)
}

func isSymmetricRecursion(left *TreeNode, right *TreeNode) bool {
	if nil == left && nil == right {
		return true
	}
	if nil == left || nil == right || left.Val != right.Val {
		return false
	}
	outterRes := isSymmetricRecursion(left.Left, right.Right)
	innerRes := isSymmetricRecursion(left.Right, right.Left)
	return outterRes && innerRes
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
func isSymmetric1(root *TreeNode) bool {
	if nil == root {
		return true
	}
	myStack := &stack{[]*TreeNode{}, 0}
	// 左右子节点都压栈
	push(myStack, root.Left)
	push(myStack, root.Right)

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
		push(myStack, right.Right)
		push(myStack, left.Right)
		push(myStack, right.Left)
	}
	return true
}
