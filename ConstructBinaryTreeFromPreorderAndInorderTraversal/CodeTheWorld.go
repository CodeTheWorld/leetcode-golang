package main

import "fmt"

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	buildTree(preorder, inorder)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
  思路：递归
  时间复杂度：O(nlog(n))
  空间复杂度：O(n)
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	length := len(preorder)
	if 0 == length {
		return nil
	}
	fmt.Println(preorder[0])
	mid := 0
	for i := 0; i < length; i++ {
		if preorder[0] == inorder[i] { // 查找根节点
			mid = i
			break
		}
	}

	return &TreeNode{preorder[0], buildTree(preorder[1:mid+1], inorder[0:mid]), buildTree(preorder[mid+1:length], inorder[mid+1:length])}
}
