package main

// https://leetcode.cn/problems/remove-element/

func removeElement(nums []int, val int) int {
	var l = 0
	var r = len(nums) - 1
	for l <= r {
		if nums[l] != val {
			l++
		} else {
			nums[l] = nums[r]
			r--
		}
	}
	return l
}
