package main

import "fmt"

func main() {
	root := &TreeNode{1, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Right.Left = &TreeNode{3, nil, nil}
	fmt.Println(inorderTraversal1(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
  思路：递归遍历
*/
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if nil != root {
		leftRes := inorderTraversal(root.Left)
		rightRes := inorderTraversal(root.Right)
		res = append(leftRes, root.Val)
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
  思路：非递归遍历，利用栈来实现
*/
func inorderTraversal1(root *TreeNode) []int {
	res := []int{}
	myStack := &stack{[]*TreeNode{}, 0}
	for nil != root || myStack.Top != 0 { // 子树根节点非空或栈非空
		for nil != root { // 子树根节点非空
			push(myStack, root)
			root = root.Left
		}
		root = pop(myStack)
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
