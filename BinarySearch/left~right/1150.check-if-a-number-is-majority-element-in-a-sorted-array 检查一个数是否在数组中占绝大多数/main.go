package main

// https://leetcode-cn.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array

func isMajorityElement(nums []int, target int) bool {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] <= target {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid
		}
	}
	// 不包含
	right := hi

	lo = 0
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] < target {
			lo = mid + 1
		} else if target <= nums[mid] {
			hi = mid
		}
	}
	// 包含
	left := lo
	return len(nums)/2 < right-left
}
