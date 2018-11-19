package main

import "fmt"

func main() {
	root := &TreeNode{10, &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{7, nil, nil}}, &TreeNode{15, nil, &TreeNode{18, nil, nil}}}
	res := rangeSumBST(root, 7, 15)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
思路：二叉搜索树，找到候选根节点，之后中序遍历
时间复杂度：O(logn + m)
空间复杂度：O(1)
*/
func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	if nil != root {
		if root.Val >= L {
			sum += rangeSumBST(root.Left, L, R)
		}
		if root.Val >= L && root.Val <= R {
			sum += root.Val
		}
		if root.Val <= R {
			sum += rangeSumBST(root.Right, L, R)
		}
	}
	return sum
}
