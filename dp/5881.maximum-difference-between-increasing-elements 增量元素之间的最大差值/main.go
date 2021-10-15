package main

// https://leetcode-cn.com/problems/maximum-difference-between-increasing-elements

func maximumDifference(nums []int) int {
	var min = 1<<63 - 1
	min = nums[0]
	l := len(nums)
	var res = -1
	for i := 1; i < l; i++ {
		if min < nums[i] {
			cur := nums[i] - min
			if res < cur {
				res = cur
			}
		} else if nums[i] < min {
			min = nums[i]
		}
	}
	return res
}
