package main

// https://leetcode.cn/problems/running-sum-of-1d-array

func runningSum(nums []int) []int {
	for i := range nums[1:] {
		nums[i+1] += nums[i]
	}
	return nums
}
