package main

import "fmt"

func main() {
	//tmpArr := []int{1, 1, 2}
	tmpArr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := removeDuplicates(tmpArr)
	fmt.Println(nums, tmpArr)
}

/**
  思路：快慢指针
  时间复杂度：O(n)
  空间复杂度：O(1)
*/
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	var i, j int
	for i, j = 0, 1; j < length; j++ { // i为慢指针，j为快指针
		if nums[i] == nums[j] {
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return i + 1
}
