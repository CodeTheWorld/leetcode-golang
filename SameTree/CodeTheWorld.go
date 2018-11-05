package main

import "fmt"

func main() {
	t1 := &TreeNode{1, nil, nil}
	t1.Left = &TreeNode{2, nil, nil}
	t1.Right = &TreeNode{1, nil, nil}
	t2 := &TreeNode{1, nil, nil}
	t2.Left = &TreeNode{1, nil, nil}
	t2.Right = &TreeNode{2, nil, nil}
	res := isSameTree(t1, t2)
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
