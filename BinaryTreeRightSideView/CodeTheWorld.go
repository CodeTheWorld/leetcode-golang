package main

import "fmt"

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}
	res := rightSideView(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
思路：类似层次遍历
时间复杂度：O(n)
空间复杂度：O(n)
*/
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if nil == root {
		return res
	}
	queue := []*TreeNode{root}
	length := 1
	for i := 0; 0 != length; i++ {
		res = append(res, queue[length-1].Val)
		for j := 0; j < length; j++ {
			if nil != queue[j].Left {
				queue = append(queue, queue[j].Left)
			}
			if nil != queue[j].Right {
				queue = append(queue, queue[j].Right)
			}
		}
		queue = queue[length:]
		length = len(queue)
	}
	return res
}
