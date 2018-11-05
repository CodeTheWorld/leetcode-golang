package main

import "fmt"

func main() {
	fmt.Println("vim-go")
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
func invertTree(root *TreeNode) *TreeNode {
	res := root
	if nil != root {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
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
  思路：非递归
  时间复杂度：O(n)
  空间复杂度：O(n)
*/
func invertTree(root *TreeNode) *TreeNode {
	res := root
	myStack := &stack{[]*TreeNode{}, 0}
	for nil != root || myStack.Top != 0 {
		for nil != root {
			push(myStack, root)
			root = root.Left
		}
		root = pop(myStack)
		root.Left, root.Right = root.Right, root.Left
		root = root.Left
	}

	return res
}
