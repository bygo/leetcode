package main

//
// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

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
