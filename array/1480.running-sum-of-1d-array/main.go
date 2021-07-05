package main

//Title: Running Sum of 1d Array
// Link: https://leetcode-cn.com/problems/running-sum-of-1d-array

func runningSum(nums []int) []int {
	for k := range nums[1:] {
		nums[k+1] += nums[k]
	}
	return nums
}
