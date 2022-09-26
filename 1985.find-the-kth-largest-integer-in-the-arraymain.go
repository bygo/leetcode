package main

import "sort"

// https://leetcode.cn/problems/find-the-kth-largest-integer-in-the-array

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		return len(nums[j]) < len(nums[i]) || len(nums[j]) == len(nums[i]) && nums[j] < nums[i]
	})
	return nums[k-1]
}
