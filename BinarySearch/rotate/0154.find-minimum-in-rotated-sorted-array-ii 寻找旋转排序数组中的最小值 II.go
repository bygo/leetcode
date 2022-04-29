package main

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[hi] < nums[mid] {
			// lo ~ mid 是旋转序列，旋转肯定偏大
			lo = mid + 1
		} else if nums[mid] < nums[hi] {
			// mid ~ hi 是正常序列
			hi = mid
		} else if nums[hi] == nums[mid] {
			// 重复元素处理 3 3 1 3
			hi--
		}
	}
	return nums[lo]
}
