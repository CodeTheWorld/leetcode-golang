package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	res := threeSumClosest(nums, target)
	fmt.Println(res)
}

/**
  思路
	1.将数组排序
	2.fix一个数
	3.双指针左右夹逼，寻找最合适的值
  时间复杂度：O(n2)
  空间复杂度：O(1)
*/
func threeSumClosest(nums []int, target int) int {
	length, closest, gap, sum := len(nums), 0, 0, 0
	sort.Ints(nums) // 排序

	for i := 0; i < length-2; i++ { // fix一个值
		for left, right := i+1, length-1; left < right; { // 双指针从两边向中间移动
			sum = nums[left] + nums[right] + nums[i]
			diff := sum - target // 计算差值diff
			if diff > 0 {
				right--
			} else if diff < 0 {
				left++
				diff *= -1
			} else {
				return sum
			}
			if diff < gap || 0 == gap {
				gap = diff
				closest = sum
			}
		}
	}

	return closest
}
