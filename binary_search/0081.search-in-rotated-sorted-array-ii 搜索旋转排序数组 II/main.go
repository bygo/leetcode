package main

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii

func search(nums []int, target int) bool {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		cur := nums[mid]
		if cur == target {
			return true
		} else if nums[l] == cur && cur == nums[r-1] {
			l++
			r--
		} else if nums[l] <= cur { // l~mid 有序 &（左段 | 右段）
			if nums[l] <= target && target < cur {
				// 范围内
				r = mid
			} else {
				l = mid + 1
			}
		} else if cur < nums[l] { // l 左段  mid～r 有序 & 右段
			if cur < target && target <= nums[r-1] {
				// 范围内
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	return false
}
