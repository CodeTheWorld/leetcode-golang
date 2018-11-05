package main

import "fmt"

func main() {
	root := &TreeNode{1, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Right.Left = &TreeNode{3, nil, nil}
	fmt.Println(postorderTraversal1(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
  思路：递归
*/
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if nil != root {
		leftRes := postorderTraversal(root.Left)
		res = append(res, leftRes...)
		rightRes := postorderTraversal(root.Right)
		res = append(res, rightRes...)
		res = append(res, root.Val)
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

func postorderTraversal1(root *TreeNode) []int {
	res := []int{}
	myStack := &stack{[]*TreeNode{}, 0}
	var lastVisit *TreeNode

	for nil != root || 0 != myStack.Top {
		for nil != root {
			push(myStack, root)
			root = root.Left
		}
		root = pop(myStack)
		if nil == root.Right || lastVisit == root.Right {
			res = append(res, root.Val)
			lastVisit = root
			root = nil
		} else {
			push(myStack, root)
			root = root.Right
		}
	}
	return res
}
