package main

import "sort"

// https://leetcode-cn.com/problems/special-array-with-x-elements-greater-than-or-equal-x

func specialArray(nums []int) int {
	sort.Ints(nums)
	numsL := len(nums)
	lo, hi := 0, numsL
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		sufL := numsL - mid
		// sufL: max -- min
		// num:  min -- max
		if sufL <= nums[mid] {
			if mid == 0 || nums[mid-1] < sufL {
				// 大于和等于都判断
				return sufL
			}
			hi = mid
		} else if nums[mid] < sufL {
			lo = mid + 1
		}
	}
	return -1
}
