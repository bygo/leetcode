package main

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	idxLeft := searchLeft(nums, target)
	idxRight := searchRight(nums, target)
	return []int{idxLeft, idxRight}
}

func searchLeft(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] < target {
			lo = mid + 1
		} else if target <= nums[mid] {
			// 左边界
			hi = mid
		}
	}
	if lo < len(nums) && nums[lo] == target {
		return lo
	}
	return -1
}

func searchRight(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] <= target {
			lo = mid + 1
		} else if target < nums[mid] {
			// 右边界
			hi = mid
		}
	}
	if 0 < lo && target == nums[lo-1] {
		return lo - 1
	}
	return -1
}
