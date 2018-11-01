package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	//nums := []int{3, 1}
	//target := 1
	res := search(nums, target)
	fmt.Println(res)
}

/**
  思路:借助二分查找的思路，通过nums[0]和target比较来区分不同的情况
  时间复杂度：O(logn)
  空间复杂度：O(1)
*/
func search(nums []int, target int) int {
	length := len(nums)

	var mid int
	for left, right := 0, length-1; left <= right; { // left和right为当前待查的数组范围
		mid = (left + right) / 2 // 中点
		if nums[mid] == target {
			return mid
		}
		if target >= nums[0] && nums[0] > nums[mid] {
			right = mid - 1
		} else if target < nums[0] && nums[0] <= nums[mid] {
			left = mid + 1
		} else {
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
