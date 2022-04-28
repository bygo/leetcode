package main

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[lo] <= nums[mid] {
			// lo ~ mid
			if nums[lo] <= target && target < nums[mid] {
				hi = mid
			} else {
				lo = mid + 1
			}
		} else if nums[mid] < nums[lo] {
			// mid ~ hi
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	}
	if nums[lo] == target {
		return lo
	}
	return -1
}

func search_(nums []int, target int) int {
	var ok int
	lo, hi := 0, len(nums)-1

	if target < nums[0] { //  A.target 是否在右半部分
		ok++
	}

	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		cur := ok
		if nums[mid] < nums[0] { // B.mid 是否在 右半部分
			cur++
		}
		if nums[mid] < target { // C.mid 在target 左
			cur++
		}

		if cur == 1 || cur == 3 { //其中1个或者3个都满足条件时，删除左边
			lo = mid + 1
		} else {
			//其中0个或者2个都满足条件时，删除右边
			hi = mid
		}
	}
	if nums[lo] == target {
		return lo
	}
	return -1
}
