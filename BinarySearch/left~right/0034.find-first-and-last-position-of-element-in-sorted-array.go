package main

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	numsL := len(nums)
	if numsL == 0 {
		return []int{-1, -1}
	}

	lo, hi := 0, numsL-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] < target {
			lo = mid + 1
		} else if target <= nums[mid] {
			hi = mid
		}
	}

	if nums[lo] != target {
		return []int{-1, -1}
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
	return []int{left, hi - 1}
}
