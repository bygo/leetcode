package main

import "sort"

// https://leetcode-cn.com/problems/special-array-with-x-elements-greater-than-or-equal-x

func specialArray(nums []int) int {
	sort.Ints(nums)
	l1 := len(nums)
	l, r := 0, l1
	for l < r {
		mid := l + (r-l)/2
		x := l1 - mid
		if x <= nums[mid] {
			if mid == 0 || nums[mid-1] < x {
				return x
			}
			r = mid
		} else if nums[mid] < x {
			l = mid + 1
		}
	}
	return -1
}
