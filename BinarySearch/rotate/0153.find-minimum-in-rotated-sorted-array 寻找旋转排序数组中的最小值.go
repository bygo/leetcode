package main

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[hi] <= nums[mid] {
			// lo ~ mid 是旋转序列，旋转肯定偏大
			lo = mid + 1
		} else if nums[mid] < nums[hi] {
			// mid ~ hi 是正常序列
			hi = mid
		}
	}
	return nums[lo]
}
