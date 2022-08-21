package main

// https://leetcode.cn/problems/check-if-a-number-is-majority-element-in-a-sorted-array

func isMajorityElement(nums []int, target int) bool {
	numsL := len(nums)
	lo, hi := 0, numsL-1 // Cause `nums[left]` might overflow
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] < target {
			lo = mid + 1
		} else if target <= nums[mid] {
			hi = mid
		}
	}
	left := lo
	hi = numsL
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] <= target {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid
		}
	}

	return nums[left] == target && numsL/2 < hi-left
}
