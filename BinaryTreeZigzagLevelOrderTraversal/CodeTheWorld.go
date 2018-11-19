package main

import "fmt"

func main() {
	//root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	res := zigzagLevelOrder1(root)
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	return zigzagLevelOrderRecursion(root, 0, [][]int{})
}

func zigzagLevelOrderRecursion(root *TreeNode, depth int, res [][]int) [][]int {
	if nil == root {
		return res
	}
	if len(res) <= depth {
		res = append(res, []int{})
	}
	if 1 == depth%2 { // 奇数行，从右向左
		res[depth] = append([]int{root.Val}, res[depth]...)
	} else { // 偶数行，从左向右
		res[depth] = append(res[depth], root.Val)
	}
	res = zigzagLevelOrderRecursion(root.Left, depth+1, res)
	res = zigzagLevelOrderRecursion(root.Right, depth+1, res)
	return res
}

/**
  思路：利用slice，迭代
  时间复杂度：O(n)
  空间复杂度：O(n)
*/
func zigzagLevelOrder1(root *TreeNode) [][]int {
	res := [][]int{}
	if nil == root {
		return res
	}
	queue := []*TreeNode{root}
	length := len(queue)
	for depth := 0; 0 != length; depth++ {
		res = append(res, []int{})
		if 0 == depth%2 { // 偶数行，从左向右
			for i := 0; i < length; i++ {
				res[depth] = append(res[depth], queue[i].Val)
				if nil != queue[i].Left {
					queue = append(queue, queue[i].Left)
				}
				if nil != queue[i].Right {
					queue = append(queue, queue[i].Right)
				}
			}
		} else { // 奇数行，从右向左
			tmpQueue := []*TreeNode{}
			for i := length - 1; i >= 0; i-- {
				res[depth] = append(res[depth], queue[i].Val)
				if nil != queue[i].Right {
					tmpQueue = append([]*TreeNode{queue[i].Right}, tmpQueue...)
				}
				if nil != queue[i].Left {
					tmpQueue = append([]*TreeNode{queue[i].Left}, tmpQueue...)
				}
			}
			queue = append(queue, tmpQueue...)
		}
		queue = queue[length:]
		length = len(queue)
	}
	return res
}
