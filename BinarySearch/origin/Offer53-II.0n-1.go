package main

// https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/

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
