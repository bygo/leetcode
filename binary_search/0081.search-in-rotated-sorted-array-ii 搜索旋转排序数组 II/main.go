package main

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii

func search(nums []int, target int) bool {
	n := len(nums)
	l, r := 0, n-1
	if n == 1 {
		return nums[0] == target
	}
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
		} else if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
