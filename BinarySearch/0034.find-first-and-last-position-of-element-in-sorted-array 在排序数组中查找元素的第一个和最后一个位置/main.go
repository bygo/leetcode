package main

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l := searchLeft(nums, target)
	r := searchRight(nums, target)
	return []int{l, r}
}

func searchLeft(nums []int, target int) int {
	var mid int
	l, r := 0, len(nums)
	for l < r {
		mid = l + (r-l)/2
		if target <= nums[mid] {
			r = mid
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	if l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}

func searchRight(nums []int, target int) int {
	var mid int
	l, r := 0, len(nums)
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid
		}
	}
	if 0 < l && target == nums[l-1] {
		return l - 1
	}
	return -1
}
