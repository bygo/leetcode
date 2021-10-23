package main

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[r] <= nums[mid] {
			l = mid + 1
		}
	}
	return nums[l]
}

// 左 去r
// 右 去r
// 左右 去r
