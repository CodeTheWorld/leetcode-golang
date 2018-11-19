package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColorMerge(nums)
	//sortColorsShell(nums)
	fmt.Println(nums)
}

/**
思路：冒泡排序
时间复杂度：O(n2)
空间复杂度：O(1)
*/
func sortColorsBubble(nums []int) {
	length := len(nums)
	for j := length; j > 0; j-- {
		for i := 0; i < j-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}

/**
思路：插入排序
时间复杂度：O(n2)
空间复杂度：O(1)
*/
func sortColorsInsertion(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		tmp := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] <= tmp {
				nums[j+1] = tmp
				break
			}
			nums[j+1] = nums[j]
		}
		if 0 > j {
			nums[0] = tmp
		}
	}
}

/**
思路：希尔排序
时间复杂度：O()
空间复杂度：O(1)
*/
func sortColorsShell(nums []int) {
	length := len(nums)
	for step := length / 2; step > 0; step /= 2 {
		for i := 0; i < step; i++ {
			for j := i; j < length; j += step {
				tmp := nums[j]
				m := j - step
				for ; m >= i; m -= step {
					if nums[m] <= tmp {
						nums[m+step] = tmp
						break
					}
					nums[m+step] = nums[m]
				}
				if m < i {
					nums[i] = tmp
				}
			}
		}
	}
}

/**
思路：选择排序
时间复杂度：O(n2)
空间复杂度：O(1)
*/
func sortColorsSelection(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
}

/**
思路：快排
时间复杂度：O(nlog(n))
空间复杂度：O(1)
*/
func sortColorsQuick(nums []int) {
	length := len(nums)
	quickSort(nums, 0, length-1)
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	tmp := nums[left]
	for i < j {
		for i < j && nums[j] >= tmp {
			j--
		}
		if i < j {
			nums[i] = nums[j]
			i++
		}
		for i < j && nums[i] <= tmp {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}
	nums[i] = tmp
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

/**
思路：归并排序
时间复杂度：O(nlogn)
空间复杂度：O(n)
*/
func sortColorMerge(nums []int) {
	mergeRecursion(nums, 0, len(nums)-1)
}

func mergeRecursion(nums []int, first int, last int) {
	mid := (first + last) / 2
	if last > first {
		mergeRecursion(nums, first, mid)
		mergeRecursion(nums, mid+1, last)
		merge(nums, first, mid, last)
	}
	return
}

func merge(nums []int, first int, mid int, last int) {
	tmpArr := []int{}
	i, j := first, mid+1
	for i <= mid && j <= last {
		if nums[i] < nums[j] {
			tmpArr = append(tmpArr, nums[i])
			i++
		} else {
			tmpArr = append(tmpArr, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmpArr = append(tmpArr, nums[i])
	}
	for ; j <= last; j++ {
		tmpArr = append(tmpArr, nums[j])
	}

	for i := first; i <= last; i++ {
		nums[i] = tmpArr[i-first]
	}
}
