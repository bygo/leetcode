package main

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	var mid, ok int
	l, r := 0, len(nums)-1

	if target < nums[0] { // 是否在右
		ok++
	}

	for l < r {
		mid = l + (r-l)/2
		cur := ok
		if nums[mid] < nums[0] { // 是否在右
			cur++
		}
		if nums[mid] < target { // 是否在目标左
			cur++
		}

		if cur == 1 || cur == 3 { //其中1个或者3个都满足条件时，删除左边
			l = mid + 1
		} else { //其中0个或者2个都满足条件时，删除右边
			r = mid
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func search(nums []int, target int) int {
	n := len(nums) - 1
	var mid int
	l, r := 0, len(nums)-1
	for l < r {
		mid = l + (r-l)/2
		if nums[0] < nums[mid] { // 左
			if nums[0] < target && target < nums[mid] { // 第一段 一定范围内
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] <= nums[0] { // 右
			if nums[mid] < target && target <= nums[n] { // 第二段 一定范围内
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
}
