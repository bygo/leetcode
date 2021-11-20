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
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		cur := nums[mid]
		if cur == target {
			return mid
		} else if nums[l] <= cur {
			if nums[l] <= target && target < cur {
				r = mid
			} else {
				l = mid + 1
			}
		} else if cur < nums[l] {
			if cur < target && target <= nums[r-1] {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	if 0 < l && nums[l-1] == target {
		return l
	}
	return -1
}
