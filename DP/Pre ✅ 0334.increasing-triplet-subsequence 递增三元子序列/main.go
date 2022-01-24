package main

// https://leetcode-cn.com/problems/increasing-triplet-subsequence

// ❓ 递增三元子序列

func increasingTriplet(nums []int) bool {
	numsL := len(nums)
	if numsL < 3 {
		return false
	}
	// 前序最小值
	minPre := make([]int, numsL)
	minPre[0] = nums[0]
	for i := 1; i < numsL; i++ {
		minPre[i] = min(minPre[i-1], nums[i])
	}

	// 后序最大值
	maxSuf := make([]int, numsL)
	maxSuf[numsL-1] = nums[numsL-1]
	for i := numsL - 2; 0 <= i; i-- {
		maxSuf[i] = max(maxSuf[i+1], nums[i])
	}

	// 判断前后
	for i := 1; i < numsL-1; i++ {
		if minPre[i-1] < nums[i] && nums[i] < maxSuf[i+1] {
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
