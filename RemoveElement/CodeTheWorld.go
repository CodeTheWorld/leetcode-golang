package main

import "fmt"

func main() {
	//nums := []int{3, 2, 2, 3}
	//val := 3
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	res := removeElement(nums, val)
	fmt.Println("nums", nums, "res", res)
}

/**
  思路：快慢指针
  时间复杂度：O(n)
  空间复杂度：O(1)
*/
func removeElement(nums []int, val int) int {
	end := 0
	for _, value := range nums {
		if val == value {
			continue
		}
		nums[end] = value
		end++
	}
	return end
}
