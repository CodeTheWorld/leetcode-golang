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

func levelOrder(root *TreeNode) [][]int {
}

func preOrder(root *TreeNode, depth int) {
	res := [][]int{}
	if nil != root {
		if depth <= len(res) {
			res[depth] = append(res[depth], root.Val)
		} else {
			res = append(res, []int{root.Val})
		}
		leftRes := preOrder(root.Left, depth++)
		rightRes := preOrder(root.Right, depth++)
	}
	return res
}
