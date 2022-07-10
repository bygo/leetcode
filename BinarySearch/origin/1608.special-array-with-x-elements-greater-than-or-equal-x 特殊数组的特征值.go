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
		if nums[mid] < sufL { // 长度过大
			lo = mid + 1
		} else if sufL <= nums[mid] {
			// nums[mid] 不可能总刚好等于 sufL
			// 所以 大于和等于都判断
			if mid == 0 || nums[mid-1] < sufL { // 前一个刚好，小于sufL
				return sufL
			}
			hi = mid
		}
	}
	return -1
}
