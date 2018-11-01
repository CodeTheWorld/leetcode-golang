package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	res := searchInsert(nums, target)
	fmt.Println(res)
}

/**
  思路：遍历
  时间复杂度：O(n)
  空间复杂度：O(1)
*/
func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if target <= v {
			return i
		}
	}
	return len(nums)
}
