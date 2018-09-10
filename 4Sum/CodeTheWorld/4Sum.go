package CodeTheWorld

import (
	"sort"
)

func FourSum(nums []int, target int) [][]int {
	var res [][]int
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := target - nums[i] - nums[j]
			for m, n := j+1, length-1; m < n; {
				tmpSum := nums[m] + nums[n]
				if tmpSum > sum {
					n--
					continue
				}
				if tmpSum < sum {
					m++
					continue
				}
				res = append(res, []int{nums[i], nums[j], nums[m], nums[n]})
				n--
				for m < n && nums[m]+nums[n] == sum {
					n--
				}
				m++
				for m < n && nums[m]+nums[n+1] == sum {
					m++
				}
			}
		}
	}
	return res
}
