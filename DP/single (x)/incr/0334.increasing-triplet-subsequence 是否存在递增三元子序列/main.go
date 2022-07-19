package main

// https://leetcode.cn/problems/increasing-triplet-subsequence

// ❓ 是否存在递增三元子序列

func increasingTriplet(nums []int) bool {
	numsL := len(nums)
	preMin := make([]int, numsL)
	preMin[0] = nums[0]
	sufMax := make([]int, numsL)
	sufMax[numsL-1] = nums[numsL-1]
	for i := 1; i < numsL; i++ {
		preMin[i] = min(preMin[i-1], nums[i])
		sufMax[numsL-i-1] = max(sufMax[numsL-i], nums[numsL-i-1])
	}

	top := numsL - 1
	for i := 1; i < top; i++ {
		if preMin[i-1] < nums[i] && nums[i] < sufMax[i+1] {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
