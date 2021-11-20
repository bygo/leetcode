package main

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii

func findMin(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[r] < nums[mid] {
			l = mid + 1
		} else if nums[r] == nums[mid] {
			r--
		}
	}
	return nums[l]
}
