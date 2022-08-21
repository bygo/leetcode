package main

// https://leetcode.cn/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] < target {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid
		} else {
			return mid
		}
	}
	return lo
}

func searchInsert(nums []int, target int) int {
	lo, hi := -1, len(nums)-1
	for lo < hi {
		mid := int(uint(lo+hi)>>1) + 1
		if nums[mid] < target {
			lo = mid
		} else if target < nums[mid] {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return lo + 1
}
