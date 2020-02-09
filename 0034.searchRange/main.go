package main

import "fmt"

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
 */

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := searchMin(nums, target)
	right := searchMax(nums, target)
	return []int{left, right}
}

func searchMin(nums []int, target int) int {
	var left, right, mid int
	right = len(nums)
	for left < right {
		mid = (left + right) / 2
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func searchMax(nums []int, target int) int {
	var left, right, mid int
	right = len(nums)
	for left < right {
		mid = (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left > 0 && nums[left-1] == target {
		return left - 1
	}
	return -1
}

func main() {

	r1 := searchRange([]int{5, 7, 7, 8, 8, 10,}, 5) //5
	//fmt.Printf("%+v", r1)
	r2 := searchRange([]int{5, 7, 7, 8, 8, 10,}, 7) //2
	//fmt.Printf("%+v", r2)

	r3 := searchRange([]int{5, 7, 7, 8, 8, 10,}, 8)  // 4
	r4 := searchRange([]int{5, 7, 7, 8, 8, 10,}, 10) //5
	r5 := searchRange([]int{5, 7, 7, 8, 8, 10,}, 6)  //-1 -1

	fmt.Printf("%+v %+v %+v %+v %+v", r1, r2, r3, r4, r5)
}
