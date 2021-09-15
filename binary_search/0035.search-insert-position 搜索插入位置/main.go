package main

// https://leetcode-cn.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			return mid
		}
	}
	return l + 1
}
