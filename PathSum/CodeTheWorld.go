package main

import "fmt"

func main() {
	//root := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, &TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, nil, &TreeNode{1, nil, nil}}}}
	//res := hasPathSum(root, 23)
	res := hasPathSum(nil, 0)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
  思路：遍历，递减
  时间复杂度：O(n)
  空间复杂度：O(1)
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if nil == root {
		return false
	}
	sum -= root.Val
	if nil == root.Left && nil == root.Right {
		return 0 == sum
	}
	res := false
	if nil != root.Left {
		leftRes := hasPathSum(root.Left, sum)
		res = res || leftRes
	}
	if nil != root.Right {
		rightRes := hasPathSum(root.Right, sum)
		res = res || rightRes
	}
	return res
}
