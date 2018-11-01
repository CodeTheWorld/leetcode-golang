package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	nums := []int{0, 0, 0}
	res := threeSum1(nums)
	fmt.Println(res)
}

/**
  思路：1.先将数组排序
  		2.fix一个值，然后用双指针法遍历右侧的子数组，寻找sum=0的情况
		3.每遇到一个值都判断和前面的值是否相等，相等就跳过（避免重复）
  时间复杂度：O(n2)
  空间复杂度：O(1)
*/
func threeSum1(nums []int) [][]int {
	length := len(nums)
	res := [][]int{}
	if length < 3 {
		return res
	}
	sort.Ints(nums) // 排序

	for i := 0; i < length-2 && nums[i] <= 0; i++ { // i为最小值，当i大于0时，不可能找到三数和为0的情况
		if i > 0 && nums[i] == nums[i-1] { // 过滤相同值
			continue
		}
		for left, right := i+1, length-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
				for nums[right] == nums[right+1] && right > left { // 过滤相同值
					right--
				}
			} else if sum < 0 {
				left++
				for nums[left] == nums[left-1] && right > left { // 过滤相同值
					left++
				}
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for nums[left] == nums[left-1] && left < right { // 过滤
					left++
				}
				for nums[right] == nums[right+1] && left < right { // 过滤
					right--
				}
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	sum := 0
	var res [][]int

	for i := 0; i < length-2; i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, length-1; j < k; {
			sum = nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			} else if sum < 0 {
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for nums[j] == nums[j-1] && j < k {
					j++
				}
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			}
		}
	}

	return res
}
