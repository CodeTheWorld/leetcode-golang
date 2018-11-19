package main

import "fmt"

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	//root := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	res := minDepth1(root)
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
  空间复杂度：O(n)
*/
func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	left := minDepth(root.Left)   // 左子树递归
	right := minDepth(root.Right) // 右子树递归
	if 0 == left {                // 左子树为空，向右延伸
		return right + 1
	} else if 0 == right { // 右子树为空，向左延伸
		return left + 1
	} else { // 左右子树都不为空，向高度小的延伸
		if left < right {
			return left + 1
		} else {
			return right + 1
		}
	}
}

/**
  思路：利用队列，迭代
  时间复杂度：O(n)
  空间复杂度：O(n)
*/
func minDepth1(root *TreeNode) int {
	depth := 0
	queue := []*TreeNode{root}
	length := len(queue)

	for length > 0 {
		for i := 0; i < length; i++ {
			if nil == queue[i] {
				continue
			}
			if nil == queue[i].Left && nil == queue[i].Right {
				return depth + 1
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
