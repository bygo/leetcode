package main

// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

func missingNumber(nums []int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if mid == nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
