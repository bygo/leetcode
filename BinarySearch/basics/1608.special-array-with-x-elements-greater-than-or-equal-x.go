package main

import "sort"

// https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x

// general
func specialArray(nums []int) int {
	sort.Ints(nums)
	numsL := len(nums)
	lo, hi := 0, numsL
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		sufL := numsL - mid
		// sufL: max -- min
		// num:  min -- max
		if nums[mid] < sufL {
			lo = mid + 1
		} else if sufL <= nums[mid] {
			if mid == 0 || nums[mid-1] < sufL { // 前一个刚好，小于sufL
				return sufL
			}
			hi = mid
		}
	}
	return -1
}

// offset
func specialArray(nums []int) int {
	sort.Ints(nums)
	numsL := len(nums)
	lo, hi := -1, numsL-1
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		sufL := numsL - mid
		if nums[mid] < sufL {
			lo = mid
		} else if sufL <= nums[mid] {
			if mid == 0 || nums[mid-1] < sufL {
				return sufL
			}
			hi = mid - 1
		}
	}
	return -1
}

// equal
func specialArray(nums []int) int {
	sort.Ints(nums)
	numsL := len(nums)
	lo, hi := 0, numsL-1
	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)
		sufL := numsL - mid
		if nums[mid] < sufL {
			lo = mid + 1
		} else if sufL <= nums[mid] {
			if mid == 0 || nums[mid-1] < sufL {
				return sufL
			}
			hi = mid - 1
		}
	}
	return -1
}
