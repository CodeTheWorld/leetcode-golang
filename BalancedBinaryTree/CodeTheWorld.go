package main

import "fmt"

func main() {
	//root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}}}
	res := isBalanced(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
  思路：递归，利用深度计算函数，当节点的左右子树高度差大于1时，直接返回-1
  时间复杂度：O(n)
  空间复杂度：O(1)
*/
func isBalanced(root *TreeNode) bool {
	depth := isBalancedRecursion(root) // 获取深度
	if depth < 0 {
		return false
	}
	return true
}

func isBalancedRecursion(root *TreeNode) int {
	if nil == root {
		return 0
	}
	leftDepth := isBalancedRecursion(root.Left)
	rightDepth := isBalancedRecursion(root.Right)
	if leftDepth < 0 || rightDepth < 0 { // 子树不平衡，那说明该树也是不平衡的
		return -1
	}
	diff := leftDepth - rightDepth
	if diff < -1 || diff > 1 { // 高度差大于1，该树不平衡
		return -1
	} else if diff >= 0 { // 高度差小于1，返回树的深度
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
