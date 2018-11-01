package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	res := combinationSum(candidates, target)
	fmt.Println(res)
}

/**
  思路：回溯算法，递归解法
  	TODO:非递归算法如何实现？
  时间复杂度：n的n次方
  空间复杂度：O(1)
*/
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return getNeed(candidates, target, 0) // 递归
}

/**
  @params candidates 数组
  @params target 目标值
  @params start 数组起始下标（防重复）
*/
func getNeed(candidates []int, target int, start int) [][]int {
	res := [][]int{}
	for i, v := range candidates {
		if i < start {
			continue
		}
		if target == v {
			res = append(res, []int{v})
		} else if v < target {
			tmpRes := getNeed(candidates, target-v, i)
			for _, value := range tmpRes {
				res = append(res, append(value, v))
			}
		} else {
			break
		}
	}
	return res
}
