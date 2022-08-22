package main

// https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/

// general
func missingNumber(nums []int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] == mid {
			lo = mid + 1
		} else if mid < nums[mid] {
			hi = mid
		}
	}
	return lo
}

// offset
func missingNumber(nums []int) int {
	lo, hi := -1, len(nums)-1
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		if nums[mid] == mid {
			lo = mid
		} else if mid < nums[mid] {
			hi = mid - 1
		}
	}
	return lo + 1
}
