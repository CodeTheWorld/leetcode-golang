package main

import "fmt"

func main() {
	root := &TreeNode{1, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Right.Left = &TreeNode{3, nil, nil}
	fmt.Println(preorderTraversal1(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
  思路：递归
*/
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if nil != root {
		res = append(res, root.Val)
		leftRes := preorderTraversal(root.Left)
		res = append(res, leftRes...)
		rightRes := preorderTraversal(root.Right)
		res = append(res, rightRes...)
	}
	return res
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
  思路：非递归，利用栈实现
*/
func preorderTraversal1(root *TreeNode) []int {
	res := []int{}
	myStack := &stack{[]*TreeNode{}, 0}
	for nil != root || 0 != myStack.Top {
		for nil != root {
			res = append(res, root.Val)
			push(myStack, root.Right)
			root = root.Left
		}
		root = pop(myStack)
	}
	return res
}
