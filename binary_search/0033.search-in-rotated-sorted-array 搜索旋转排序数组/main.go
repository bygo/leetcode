package main

// search in rotated sorter array
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	var left, mid, ok int
	right := len(nums) - 1
	for left < right {
		mid = left + (right-left)/2
		ok = 0
		if target < nums[0] {
			ok++
		}
		if nums[mid] < nums[0] {
			ok++
		}
		if nums[mid] < target {
			ok++
		}
		if ok == 1 || ok == 3 { //其中1个或者3个都满足条件时，删除左边
			left = mid + 1
		} else { //其中0个或者2个都满足条件时，删除右边
			right = mid
		}
	}
	if left == right && nums[left] == target {
		return left
	}
	return -1
}
