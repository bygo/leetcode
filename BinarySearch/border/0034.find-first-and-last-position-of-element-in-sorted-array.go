package main

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	numsL := len(nums)
	if numsL == 0 {
		return []int{-1, -1}
	}

	lo, hi := 0, numsL-1 // TODO
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

func searchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums)
	if hi == 0 {
		return []int{-1, -1}
	}
	var left, right int
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] < target {
			lo = mid + 1
		} else if target <= nums[mid] {
			hi = mid
		}
	}
	left = lo
	hi = len(nums)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] <= target {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid
		}
	}
	right = hi - 1
	if -1 == left || nums[left] != target {
		return []int{-1, -1}
	}
	return []int{left, right}
}
