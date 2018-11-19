package main

import "fmt"

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	buildTree(inorder, postorder)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
思路：递归
时间复杂度：O(n*log(n))
空间复杂度：O(n)
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	length := len(inorder)
	if 0 == length {
		return nil
	}
	mid := 0
	for i := 0; i < length; i++ {
		if inorder[i] == postorder[length-1] {
			mid = i
			break
		}
	}
	fmt.Println(postorder[length-1])

	return &TreeNode{postorder[length-1], buildTree(inorder[0:mid], postorder[0:mid]), buildTree(inorder[mid+1:length], postorder[mid:length-1])}
}
