package main

// https://leetcode-cn.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array

func isMajorityElement(nums []int, target int) bool {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid
		}
	}
	right := r

	l, r = 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if target <= nums[mid] {
			r = mid
		}
	}
	left := l
	return len(nums)/2 < right-left
}
