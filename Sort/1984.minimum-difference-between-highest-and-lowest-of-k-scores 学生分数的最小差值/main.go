package main

import "sort"

// https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores

// ❓ 学生分数的最小差值

func minimumDifference(nums []int, k int) int {
	numsL := len(nums)
	sort.Ints(nums)
	var diffMin = 1<<63 - 1
	k -= 1
	for i := k; i < numsL; i++ {
		diffMin = min(diffMin, nums[i]-nums[i-k])
	}
	return diffMin
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
