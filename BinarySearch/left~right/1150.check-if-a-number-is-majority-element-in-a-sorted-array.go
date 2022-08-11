package main

// https://leetcode.cn/problems/check-if-a-number-is-majority-element-in-a-sorted-array

func isMajorityElement(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] < target {
			lo = mid + 1
		} else if target <= nums[mid] {
			hi = mid
		}
	}
	// 不包含
	left := lo

	hi = len(nums) - 1
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		if nums[mid] <= target {
			lo = mid
		} else if target < nums[mid] {
			hi = mid - 1
		}
	}
	// 包含
	right := hi
	return nums[left] == target && len(nums)/2 < right-left+1
}

func isMajorityElement(nums []int, target int) bool {
	numsL := len(nums)
	lo, hi := 0, numsL-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] < target {
			lo = mid + 1
		} else if target <= nums[mid] {
			hi = mid
		}
	}
	left := lo
	hi = numsL - 1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] <= target {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid
		}
	}

	return nums[left] == target && numsL/2 < hi-left+1
}
