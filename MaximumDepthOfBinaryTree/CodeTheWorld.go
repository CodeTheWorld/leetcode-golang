package main

import "fmt"

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	res := maxDepth1(root)
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
func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

/**
  思路：利用队列，迭代
  时间复杂度：O(n)
  空间复杂度：O(n)
*/
func maxDepth1(root *TreeNode) int {
	depth := 0
	queue := []*TreeNode{root}
	length := len(queue)

	for length > 0 {
		for i := 0; i < length; i++ {
			if nil == queue[i] {
				continue
			}
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
		queue = queue[length:]
		length = len(queue)
		if length > 0 {
			depth++
		}
	}
	return depth
}
