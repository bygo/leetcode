package main

// https://leetcode.cn/problems/maximum-difference-between-increasing-elements

func maximumDifference(nums []int) int {
	min := nums[0]
	numsL := len(nums)
	var numDiffMax = -1
	for i := 1; i < numsL; i++ {
		if nums[i] < min {
			min = nums[i]
		} else if min < nums[i] {
			numDiff := nums[i] - min
			if numDiffMax < numDiff {
				numDiffMax = numDiff
			}
		}
	}
	return numDiffMax
}
