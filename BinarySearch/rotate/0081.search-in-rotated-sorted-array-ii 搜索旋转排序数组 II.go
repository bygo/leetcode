package main

// https://leetcode.cn/problems/search-in-rotated-sorted-array-ii

func search(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if nums[mid] == target {
			return true
		} else if nums[lo] == nums[mid] && nums[mid] == nums[hi] {
			lo++
			hi--
		} else if nums[lo] <= nums[mid] { // l ~ mid 有序 &（左段 | 右段）
			if nums[lo] <= target && target < nums[mid] {
				// 范围内
				hi = mid
			} else {
				lo = mid + 1
			}
		} else if nums[mid] < nums[lo] { // l 左段  mid～r 有序 & 右段
			if nums[mid] < target && target <= nums[hi] {
				// 范围内
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	}
	return nums[lo] == target
}
